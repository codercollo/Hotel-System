// Package apierror defines structured error types used throughout the backend.
// Every domain error maps to exactly one HTTP status code and one machine-readable code.
// Handlers should return APIError values; the response package renders them.
package apierror

import (
	"errors"
	"fmt"
	"net/http"
)

// Code is a machine-readable error identifier. Clients key on this value.
type Code string

const (
	// Generic
	CodeInternal       Code = "INTERNAL_ERROR"
	CodeNotFound       Code = "NOT_FOUND"
	CodeConflict       Code = "CONFLICT"
	CodeBadRequest     Code = "BAD_REQUEST"
	CodeUnprocessable  Code = "UNPROCESSABLE_ENTITY"
	CodeUnauthorized   Code = "UNAUTHORIZED"
	CodeForbidden      Code = "FORBIDDEN"
	CodeTooManyReqs    Code = "TOO_MANY_REQUESTS"
	CodeServiceUnavail Code = "SERVICE_UNAVAILABLE"

	// Validation
	CodeValidation Code = "VALIDATION_ERROR"

	// Auth domain
	CodeInvalidCredentials Code = "INVALID_CREDENTIALS"
	CodeTokenExpired       Code = "TOKEN_EXPIRED"
	CodeTokenInvalid       Code = "TOKEN_INVALID"

	// Resource-specific (domain-agnostic)
	CodeItemNotFound  Code = "ITEM_NOT_FOUND"
	CodeOrderNotFound Code = "ORDER_NOT_FOUND"
	CodePaymentFailed Code = "PAYMENT_FAILED"
	CodeUserNotFound  Code = "USER_NOT_FOUND"
	CodeUserExists    Code = "USER_ALREADY_EXISTS"
	CodeRoleNotFound  Code = "ROLE_NOT_FOUND"
	CodeUploadFailed  Code = "UPLOAD_FAILED"
)

// APIError is the canonical error type for domain and handler errors.
// It carries an HTTP status, machine code, human message, and optional detail.
type APIError struct {
	Status  int    `json:"-"`
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
	err     error
}

func (e *APIError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap enables errors.Is / errors.As on the wrapped cause.
func (e *APIError) Unwrap() error { return e.err }

// WithCause attaches a causal error (logged but not exposed to clients).
func (e *APIError) WithCause(err error) *APIError {
	e.err = err
	return e
}

// WithDetail attaches structured detail for validation errors, etc.
func (e *APIError) WithDetail(d any) *APIError {
	e.Detail = d
	return e
}

// ─── Constructors ─────────────────────────────────────────────────────────────

func New(status int, code Code, message string) *APIError {
	return &APIError{Status: status, Code: code, Message: message}
}

func Internal(err error) *APIError {
	return New(http.StatusInternalServerError, CodeInternal, "an unexpected error occurred").WithCause(err)
}

func NotFound(code Code, message string) *APIError {
	return New(http.StatusNotFound, code, message)
}

func BadRequest(message string) *APIError {
	return New(http.StatusBadRequest, CodeBadRequest, message)
}

func Validation(detail any) *APIError {
	return New(http.StatusUnprocessableEntity, CodeValidation, "validation failed").WithDetail(detail)
}

func Unauthorized(message string) *APIError {
	return New(http.StatusUnauthorized, CodeUnauthorized, message)
}

func Forbidden(message string) *APIError {
	return New(http.StatusForbidden, CodeForbidden, message)
}

func Conflict(code Code, message string) *APIError {
	return New(http.StatusConflict, code, message)
}

func TooManyRequests() *APIError {
	return New(http.StatusTooManyRequests, CodeTooManyReqs, "rate limit exceeded")
}

// ─── Sentinel errors for common cases ────────────────────────────────────────

var (
	ErrNotFound     = NotFound(CodeNotFound, "resource not found")
	ErrUnauthorized = Unauthorized("authentication required")
	ErrForbidden    = Forbidden("insufficient permissions")
	ErrTokenExpired = New(http.StatusUnauthorized, CodeTokenExpired, "token has expired")
	ErrTokenInvalid = New(http.StatusUnauthorized, CodeTokenInvalid, "token is invalid")
	ErrInvalidCreds = New(http.StatusUnauthorized, CodeInvalidCredentials, "invalid credentials")
	ErrInternal     = New(http.StatusInternalServerError, CodeInternal, "an unexpected error occurred")
	ErrUserExists   = Conflict(CodeUserExists, "a user with this email already exists")
)

// ─── Type utilities ───────────────────────────────────────────────────────────

// IsAPIError reports whether err is (or wraps) an *APIError and extracts it.
func IsAPIError(err error) (*APIError, bool) {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}

// NewValidation creates a 422 Unprocessable Entity error with per-field messages.
func NewValidation(fields map[string]string) *APIError {
	return &APIError{
		Status:  http.StatusUnprocessableEntity,
		Code:    CodeValidation,
		Message: "one or more fields failed validation",
		Detail:  fields,
	}
}
