package clashy

import "fmt"

type ClashOfClansException struct {
	Message string
}

func (e *ClashOfClansException) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

type HTTPException struct {
	Status  int
	Reason  string
	Message string
	Body    []byte
}

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

type InvalidArgument struct{ *HTTPException }
type InvalidCredentials struct{ *HTTPException }
type Forbidden struct{ *HTTPException }
type PrivateWarLog struct{ *HTTPException }
type NotFound struct{ *HTTPException }
type Maintenance struct{ *HTTPException }
type GatewayError struct{ *HTTPException }
