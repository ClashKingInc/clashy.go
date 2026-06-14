package clashy

import "fmt"

// ClashOfClansException is the base package error type.
type ClashOfClansException struct {
	// Message is the human-readable error message.
	Message string
}

// Error implements the error interface.
func (e *ClashOfClansException) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

// HTTPException captures a non-successful API response.
//
// The typed HTTP errors in this package embed HTTPException so callers can
// match either the specific type or inspect shared status, reason, message, and
// body fields.
type HTTPException struct {
	// Status is the HTTP status code returned by the API. It can be zero for
	// transport failures mapped to GatewayError.
	Status int
	// Reason is the API reason string when one was provided.
	Reason string
	// Message is the API message string when one was provided.
	Message string
	// Body is the raw response body retained for debugging.
	Body []byte
}

// Error implements the error interface.
func (e *HTTPException) Error() string {
	if e == nil {
		return ""
	}
	if e.Message != "" {
		return fmt.Sprintf("%s (status code: %d): %s", e.Reason, e.Status, e.Message)
	}
	return fmt.Sprintf("%s (status code: %d)", e.Reason, e.Status)
}

func newHTTPException(status int, reason, message string, body []byte) *HTTPException {
	if reason == "" {
		reason = "Unknown"
	}
	return &HTTPException{
		Status:  status,
		Reason:  reason,
		Message: message,
		Body:    body,
	}
}

// InvalidArgument represents a 400 response from the API.
type InvalidArgument struct{ *HTTPException }

// InvalidCredentials represents a developer-site authentication failure.
type InvalidCredentials struct{ *HTTPException }

// Forbidden represents a 403 response from the API.
type Forbidden struct{ *HTTPException }

// PrivateWarLog represents the private-war-log 403 response.
type PrivateWarLog struct{ *HTTPException }

// NotFound represents a 404 response from the API.
type NotFound struct{ *HTTPException }

// Maintenance represents a 503 maintenance response from the API.
type Maintenance struct{ *HTTPException }

// GatewayError represents transport failures and 5xx gateway responses.
type GatewayError struct{ *HTTPException }
