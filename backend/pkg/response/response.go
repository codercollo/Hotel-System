// Package response provides standardised JSON envelope helpers for HTTP handlers.
// All API responses use one of two shapes:
//
//	Success: { "success": true,  "data": <T>, "meta": <pagination | null> }
//	Error:   { "success": false, "error": { "code": "...", "message": "...", "detail": ... } }
package response

import (
	"encoding/json"
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
)

// envelope is the outer JSON wrapper for all responses.
type envelope struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Meta    any  `json:"meta,omitempty"`
	Error   any  `json:"error,omitempty"`
}

type errorBody struct {
	Code    apierror.Code `json:"code"`
	Message string        `json:"message"`
	Detail  any           `json:"detail,omitempty"`
}

// JSON writes a success response with HTTP 200.
func JSON(w http.ResponseWriter, data any) {
	JSONStatus(w, http.StatusOK, data, nil)
}

// JSONCreated writes a success response with HTTP 201.
func JSONCreated(w http.ResponseWriter, data any) {
	JSONStatus(w, http.StatusCreated, data, nil)
}

// JSONPaginated writes a success response with pagination metadata.
func JSONPaginated(w http.ResponseWriter, data any, meta any) {
	JSONStatus(w, http.StatusOK, data, meta)
}

// NoContent writes HTTP 204 with an empty body.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// JSONStatus writes a success envelope with a custom status code.
func JSONStatus(w http.ResponseWriter, status int, data any, meta any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// Error writes an error envelope derived from an *apierror.APIError.
// Falls back to 500 Internal Server Error for unrecognised error types.
func Error(w http.ResponseWriter, err error) {
	ae, ok := apierror.IsAPIError(err)
	if !ok {
		ae = apierror.Internal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ae.Status)
	_ = json.NewEncoder(w).Encode(envelope{
		Success: false,
		Error: errorBody{
			Code:    ae.Code,
			Message: ae.Message,
			Detail:  ae.Detail,
		},
	})
}

// ErrorStatus writes an error envelope with an explicit status code and message.
// Prefer returning apierror types from handlers and using Error() instead.
func ErrorStatus(w http.ResponseWriter, status int, code apierror.Code, message string) {
	Error(w, apierror.New(status, code, message))
}
