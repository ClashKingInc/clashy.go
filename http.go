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

type cachedResponse struct {
	Body      []byte
	Status    int
	ExpiresAt time.Time
}

type HTTPClient struct {
	config ClientConfig
	client *http.Client

	mu    sync.RWMutex
	cache map[string]cachedResponse
	keys  []string
	next  int
}

func NewHTTPClient(cfg ClientConfig) *HTTPClient {
	return &HTTPClient{
		config: cfg,
		client: &http.Client{Timeout: cfg.Timeout},
		cache:  make(map[string]cachedResponse),
	}
}

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

	responseBody, err := readResponseBody(resp)
	if err != nil {
		return nil, 0, 0, err
	}

	retry := int(cacheExpiry(resp.Header.Get("Cache-Control")).Seconds())
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if method == http.MethodGet && options.UpdateCache && retry > 0 {
			h.mu.Lock()
			h.cache[fullURL] = cachedResponse{Body: append([]byte(nil), responseBody...), Status: resp.StatusCode, ExpiresAt: time.Now().Add(time.Duration(retry) * time.Second)}
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

type RequestOptions struct {
	LookupCache bool
	UpdateCache bool
	SkipAuth    bool
}

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
	payloadBody, _ := io.ReadAll(resp.Body)
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

func readResponseBody(resp *http.Response) ([]byte, error) {
	reader, err := decodedBodyReader(resp)
	if err != nil {
		return nil, err
	}
	if reader == resp.Body {
		return io.ReadAll(reader)
	}
	defer reader.Close()
	return io.ReadAll(reader)
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
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("developer-site request %s failed: %d", path, resp.StatusCode)
	}
	if out == nil {
		return nil
	}
	return json.Unmarshal(body, out)
}
