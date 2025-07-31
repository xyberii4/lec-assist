package tlservice

import "fmt"

const (
	ErrCodeValidationFailed = "validation_failed"
	ErrCodeInvalidArg       = "invalid_argument"
	ErrCodeConflict         = "conflict"
	ErrCodeExternalAPIError = "external_api_error"

	ErrCodeRateLimitExceeded = "rate_limit_exceeded"
)

type ServiceError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *ServiceError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("service error [%s]: %s (underlying: %v)", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("service error [%s]: %s", e.Code, e.Message)
}

func (e *ServiceError) Unwrap() error {
	return e.Err
}

func NewServiceError(code, message string, err error) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
