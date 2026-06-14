package clashy

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const maxResponseBytes = 32 << 20

type cachedResponse struct {
	Body      []byte
	Status    int
	ExpiresAt time.Time
}

// HTTPClient performs low-level Clash API and developer-site HTTP requests.
//
// Most callers should use Client instead. HTTPClient is exported so advanced
// integrations can build compatible request flows while reusing token rotation,
// throttling, compression handling, cache storage, and typed error mapping.
type HTTPClient struct {
	config ClientConfig
	client *http.Client
	limit  *requestLimiter

	mu         sync.RWMutex
	cache      map[string]cachedResponse
	cacheOrder []string
	keys       []string
	next       int
}

// NewHTTPClient constructs an HTTPClient from cfg.
func NewHTTPClient(cfg ClientConfig) *HTTPClient {
	return &HTTPClient{
		config: cfg,
		client: &http.Client{Timeout: cfg.Timeout},
		limit:  newRequestLimiter(cfg.ThrottleLimit),
		cache:  make(map[string]cachedResponse),
	}
}

// SetTokens replaces the API tokens used for Authorization headers.
//
// Tokens are rotated one per request. Passing no tokens clears authentication.
func (h *HTTPClient) SetTokens(tokens ...string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.keys = append([]string(nil), tokens...)
	h.next = 0
}

func (h *HTTPClient) token() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.keys) == 0 {
		return ""
	}
	token := h.keys[h.next%len(h.keys)]
	h.next++
	return token
}

// Do sends one HTTP request and returns the response body, status code, retry
// cache duration in seconds, and error.
//
// Non-2xx API responses are converted into the package's typed HTTP errors.
// Successful GET responses can be read from or written to the in-memory cache
// depending on RequestOptions.
func (h *HTTPClient) Do(ctx context.Context, method, fullURL string, body any, options RequestOptions) ([]byte, int, int, error) {
	ctx = contextOrBackground(ctx)
	if method == http.MethodGet && options.LookupCache {
		h.mu.RLock()
		cached, ok := h.cache[fullURL]
		h.mu.RUnlock()
		if ok && time.Now().Before(cached.ExpiresAt) {
			return append([]byte(nil), cached.Body...), cached.Status, int(time.Until(cached.ExpiresAt).Seconds()), nil
		}
	}
	if !rateLimitDisabled(ctx) {
		release, err := h.limit.Acquire(ctx)
		if err != nil {
			return nil, 0, 0, err
		}
		defer release()
	}

	var reader io.Reader
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			return nil, 0, 0, err
		}
		reader = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, reader)
	if err != nil {
		return nil, 0, 0, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	if h.config.UserAgent != "" {
		req.Header.Set("User-Agent", h.config.UserAgent)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token := h.token(); token != "" && !options.SkipAuth {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, 0, 0, &GatewayError{newHTTPException(0, "Gateway Error", err.Error(), nil)}
	}
	defer resp.Body.Close()

	responseBody, err := h.readResponseBody(resp)
	if err != nil {
		return nil, 0, 0, err
	}

	retry := int(cacheExpiry(resp.Header.Get("Cache-Control")).Seconds())
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if method == http.MethodGet && options.UpdateCache && retry > 0 {
			h.mu.Lock()
			h.cacheResponse(fullURL, cachedResponse{Body: append([]byte(nil), responseBody...), Status: resp.StatusCode, ExpiresAt: time.Now().Add(time.Duration(retry) * time.Second)})
			h.mu.Unlock()
		}
		return responseBody, resp.StatusCode, retry, nil
	}

	var apiErr struct {
		Reason  string `json:"reason"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(responseBody, &apiErr)

	switch resp.StatusCode {
	case 400:
		return nil, resp.StatusCode, retry, &InvalidArgument{newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)}
	case 403:
		reason := apiErr.Reason
		if strings.EqualFold(reason, "accessDenied.privateWarLog") {
			return nil, resp.StatusCode, retry, &PrivateWarLog{newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)}
		}
		return nil, resp.StatusCode, retry, &Forbidden{newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)}
	case 404:
		return nil, resp.StatusCode, retry, &NotFound{newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)}
	case 503:
		return nil, resp.StatusCode, retry, &Maintenance{newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)}
	case 500, 502, 504:
		return nil, resp.StatusCode, retry, &GatewayError{newHTTPException(resp.StatusCode, "Gateway Error", string(responseBody), responseBody)}
	default:
		return nil, resp.StatusCode, retry, newHTTPException(resp.StatusCode, apiErr.Reason, apiErr.Message, responseBody)
	}
}

// RequestOptions controls per-request behavior for HTTPClient.Do.
type RequestOptions struct {
	// LookupCache allows a GET request to return a fresh cached response.
	LookupCache bool
	// UpdateCache allows a successful GET request to store or replace a cached
	// response.
	UpdateCache bool
	// SkipAuth prevents Do from adding an Authorization header.
	SkipAuth bool
}

// LoginDeveloper authenticates against the Clash developer site and configures
// API tokens for subsequent Clash API requests.
//
// The method reuses matching keys for the configured IP and key name when
// possible, creating more keys until ClientConfig.KeyCount is satisfied.
func (h *HTTPClient) LoginDeveloper(ctx context.Context, email, password string) error {
	ctx = contextOrBackground(ctx)
	payload := map[string]string{"email": email, "password": password}
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	loginURL := normalizeAPIBase(h.config.DeveloperBaseURL) + "/api/login"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if h.config.DeveloperUserAgent != "" {
		req.Header.Set("User-Agent", h.config.DeveloperUserAgent)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	payloadBody, err := readLimited(resp.Body, maxResponseBytes)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusForbidden {
		return &InvalidCredentials{newHTTPException(resp.StatusCode, "invalid credentials", "developer-site login failed", payloadBody)}
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("developer-site login failed with %d", resp.StatusCode)
	}

	var loginResponse struct {
		TemporaryAPIToken string `json:"temporaryAPIToken"`
	}
	if err := json.Unmarshal(payloadBody, &loginResponse); err != nil {
		return err
	}
	ip := h.config.IP
	if ip == "" {
		ip = jwtIP(loginResponse.TemporaryAPIToken)
	}

	type devKey struct {
		ID         int      `json:"id"`
		Name       string   `json:"name"`
		CIDRRanges []string `json:"cidrRanges"`
		Key        string   `json:"key"`
	}
	var keysResp struct {
		Keys []devKey `json:"keys"`
	}
	if err := h.developerJSON(ctx, "/api/apikey/list", nil, &keysResp); err != nil {
		return err
	}

	var tokens []string
	for _, key := range keysResp.Keys {
		if key.Name != h.config.KeyNames {
			continue
		}
		for _, cidr := range key.CIDRRanges {
			if strings.HasPrefix(cidr, ip) {
				tokens = append(tokens, key.Key)
				break
			}
		}
	}
	for len(tokens) < h.config.KeyCount && len(keysResp.Keys) < 10 {
		var createResp struct {
			Key devKey `json:"key"`
		}
		body := map[string]any{
			"name":        h.config.KeyNames,
			"description": "Created by clashy.go",
			"cidrRanges":  []string{ip},
			"scopes":      []string{"clash"},
		}
		if err := h.developerJSON(ctx, "/api/apikey/create", body, &createResp); err != nil {
			return err
		}
		tokens = append(tokens, createResp.Key.Key)
		keysResp.Keys = append(keysResp.Keys, createResp.Key)
	}
	h.SetTokens(tokens...)
	return nil
}

func (h *HTTPClient) readResponseBody(resp *http.Response) ([]byte, error) {
	reader, err := decodedBodyReader(resp)
	if err != nil {
		return nil, err
	}
	if reader == resp.Body {
		return readLimited(reader, maxResponseBytes)
	}
	defer reader.Close()
	return readLimited(reader, maxResponseBytes)
}

func decodedBodyReader(resp *http.Response) (io.ReadCloser, error) {
	switch strings.ToLower(strings.TrimSpace(resp.Header.Get("Content-Encoding"))) {
	case "", "identity":
		return resp.Body, nil
	case "gzip":
		return gzip.NewReader(resp.Body)
	case "deflate":
		return zlib.NewReader(resp.Body)
	default:
		return resp.Body, nil
	}
}

func (h *HTTPClient) developerJSON(ctx context.Context, path string, reqBody any, out any) error {
	var reader io.Reader
	if reqBody != nil {
		payload, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, normalizeAPIBase(h.config.DeveloperBaseURL)+path, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := readLimited(resp.Body, maxResponseBytes)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("developer-site request %s failed: %d", path, resp.StatusCode)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(body, out)
}

func readLimited(reader io.Reader, maxBytes int64) ([]byte, error) {
	if maxBytes <= 0 {
		return io.ReadAll(reader)
	}
	body, err := io.ReadAll(io.LimitReader(reader, maxBytes+1))
	if err != nil {
		return nil, err
	}
	if int64(len(body)) > maxBytes {
		return nil, fmt.Errorf("response body exceeds %d bytes", maxBytes)
	}
	return body, nil
}

func (h *HTTPClient) enforceCacheMaxSize() {
	if h.config.CacheMaxSize <= 0 {
		clear(h.cache)
		h.cacheOrder = h.cacheOrder[:0]
		return
	}
	for len(h.cache) > h.config.CacheMaxSize && len(h.cacheOrder) > 0 {
		oldest := h.cacheOrder[0]
		h.cacheOrder = h.cacheOrder[1:]
		if _, ok := h.cache[oldest]; !ok {
			continue
		}
		delete(h.cache, oldest)
	}
}

func (h *HTTPClient) cacheResponse(fullURL string, response cachedResponse) {
	if h.config.CacheMaxSize <= 0 {
		return
	}
	if _, exists := h.cache[fullURL]; !exists {
		h.cacheOrder = append(h.cacheOrder, fullURL)
	}
	h.cache[fullURL] = response
	h.enforceCacheMaxSize()
}
