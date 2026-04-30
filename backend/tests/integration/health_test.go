// Package integration contains end-to-end tests that require live infrastructure.
// Run with: go test ./tests/integration/... -tags integration
//
// Requires DATABASE_URL env var pointing at a real (test) database.
//go:build integration

package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/platform/backend/config"
	"github.com/platform/backend/database"
	"github.com/platform/backend/middleware"
	"github.com/platform/backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthEndpoint(t *testing.T) {
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("DATABASE_URL not set — skipping integration test")
	}

	cfg, err := config.Load()
	require.NoError(t, err)

	logger.Init("error", "json") // suppress output during tests

	pool, err := database.Init(t.Context(), cfg.Database)
	require.NoError(t, err)
	defer pool.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/api/health", func(w http.ResponseWriter, req *http.Request) {
		if err := database.HealthCheck(req.Context(), pool); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
	})

	srv := httptest.NewServer(r)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/health")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]any
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.Equal(t, "ok", body["status"])
}

func TestPingEndpoint(t *testing.T) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Get("/api/v1/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true,"data":"pong"}`))
	})

	srv := httptest.NewServer(r)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/v1/ping")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	// Request-ID must be injected
	assert.NotEmpty(t, resp.Header.Get("X-Request-ID"))
}
