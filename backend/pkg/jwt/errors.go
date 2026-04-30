package jwt

import "github.com/codercollo/hotel-system/backend/pkg/apierror"

// Exported sentinel errors for middleware consumption.
var (
	ErrMissingToken   = apierror.Unauthorized("authorization token is required")
	ErrWrongTokenType = apierror.Unauthorized("wrong token type")
)
