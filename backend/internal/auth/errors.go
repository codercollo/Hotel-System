package auth

import (
	"net/http"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
)

var (
	ErrSessionNotFound    = apierror.New(http.StatusUnauthorized, apierror.CodeUnauthorized, "session not found")
	ErrInvalidCredentials = apierror.New(http.StatusUnauthorized, apierror.CodeUnauthorized, "invalid credentials")
	ErrSessionBlocked     = apierror.New(http.StatusUnauthorized, apierror.CodeUnauthorized, "session has been revoked")
	ErrSessionExpired     = apierror.New(http.StatusUnauthorized, apierror.CodeUnauthorized, "session has expired")
	ErrAccountInactive    = apierror.New(http.StatusForbidden, apierror.CodeForbidden, "account is inactive")
)
