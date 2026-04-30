package config_test

import (
	"os"
	"testing"

	"github.com/codercollo/hotel-system/backend/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad_Defaults(t *testing.T) {
	// Provide the only truly required values.
	t.Setenv("DATABASE_URL", "postgres://platform:platform@localhost:5432/platform_test?sslmode=disable")
	t.Setenv("JWT_SECRET", "test-secret-that-is-long-enough-for-validation")

	cfg, err := config.Load()
	require.NoError(t, err)

	assert.Equal(t, "development", cfg.App.Env)
	assert.Equal(t, 8080, cfg.App.Port)
	assert.Equal(t, ":8080", cfg.Addr())
	assert.True(t, cfg.IsDevelopment())
	assert.False(t, cfg.IsProduction())
}

func TestLoad_MissingDatabaseURL(t *testing.T) {
	os.Unsetenv("DATABASE_URL")
	t.Setenv("JWT_SECRET", "test-secret-that-is-long-enough-for-validation")

	_, err := config.Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DATABASE_URL")
}

func TestLoad_ShortJWTSecret(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://localhost/test")
	t.Setenv("JWT_SECRET", "short")

	_, err := config.Load()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "JWT_SECRET")
}

func TestLoad_CORSOriginsParsed(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://localhost/test")
	t.Setenv("JWT_SECRET", "test-secret-that-is-long-enough-for-validation")
	t.Setenv("CORS_ORIGINS", "http://localhost:3000,https://app.example.com")

	cfg, err := config.Load()
	require.NoError(t, err)
	assert.Len(t, cfg.CORS.Origins, 2)
	assert.Contains(t, cfg.CORS.Origins, "http://localhost:3000")
	assert.Contains(t, cfg.CORS.Origins, "https://app.example.com")
}
