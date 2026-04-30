package response

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/codercollo/hotel-system/backend/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type healthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	GoVersion string            `json:"go_version"`
	Checks    map[string]string `json:"checks"`
}

func HealthHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		checks := map[string]string{}
		overall := "ok"

		// Ensure your database package has a HealthCheck function
		if err := database.HealthCheck(r.Context(), pool); err != nil {
			checks["database"] = "unhealthy: " + err.Error()
			overall = "degraded"
		} else {
			checks["database"] = "ok"
		}

		status := http.StatusOK
		if overall != "ok" {
			status = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(healthResponse{
			Status:    overall,
			Timestamp: time.Now().UTC(),
			GoVersion: runtime.Version(),
			Checks:    checks,
		})
	}
}
