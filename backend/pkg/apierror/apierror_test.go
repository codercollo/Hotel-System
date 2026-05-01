package apierror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	e := apierror.New(http.StatusNotFound, apierror.CodeNotFound, "not found")
	assert.Equal(t, http.StatusNotFound, e.Status)
	assert.Equal(t, apierror.CodeNotFound, e.Code)
	assert.Equal(t, "not found", e.Message)
}

func TestWithCause(t *testing.T) {
	cause := errors.New("db error")
	e := apierror.Internal(cause)
	assert.ErrorIs(t, e, cause)
	assert.Contains(t, e.Error(), "db error")
}

func TestIsAPIError(t *testing.T) {
	e := apierror.NotFound(apierror.CodeNotFound, "missing")
	ae, ok := apierror.IsAPIError(e)
	assert.True(t, ok)
	assert.Equal(t, http.StatusNotFound, ae.Status)
}

func TestIsAPIError_NonAPIError(t *testing.T) {
	_, ok := apierror.IsAPIError(errors.New("plain error"))
	assert.False(t, ok)
}

func TestValidation(t *testing.T) {
	detail := []map[string]string{{"field": "email", "message": "required"}}
	e := apierror.Validation(detail)
	assert.Equal(t, http.StatusUnprocessableEntity, e.Status)
	assert.Equal(t, apierror.CodeValidation, e.Code)
	assert.NotNil(t, e.Detail)
}
