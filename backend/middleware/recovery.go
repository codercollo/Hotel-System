package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/logger"
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

// Recovery catches panics, logs the stack trace, and returns HTTP 500.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err := fmt.Errorf("panic: %v\n%s", rec, debug.Stack())
				logger.Error(err, "panic recovered")
				response.Error(w, apierror.Internal(err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
