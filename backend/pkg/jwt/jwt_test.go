package jwt_test

import (
	"testing"
	"time"

	pkgjwt "github.com/codercollo/hotel-system/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newManager() *pkgjwt.Manager {
	return pkgjwt.NewManager(
		"test-secret-that-is-long-enough-ok",
		15*time.Minute,
		7*24*time.Hour,
	)
}

func TestSignAndParseAccess(t *testing.T) {
	m := newManager()

	token, err := m.SignAccess("user-1", "user@example.com", "user")
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := m.Parse(token)
	require.NoError(t, err)

	assert.Equal(t, "user-1", claims.UserID)
	assert.Equal(t, "user@example.com", claims.Email)
	assert.Equal(t, "user", claims.Role)
	assert.Equal(t, pkgjwt.TokenTypeAccess, claims.TokenType)
}

func TestSignAndParseRefresh(t *testing.T) {
	m := newManager()

	token, err := m.SignRefresh("user-2", "admin@example.com", "admin")
	require.NoError(t, err)

	claims, err := m.ParseRefresh(token)
	require.NoError(t, err)
	assert.Equal(t, pkgjwt.TokenTypeRefresh, claims.TokenType)
}

func TestParseRefresh_RejectsAccessToken(t *testing.T) {
	m := newManager()

	token, _ := m.SignAccess("user-1", "x@x.com", "user")
	_, err := m.ParseRefresh(token)
	assert.Error(t, err)
}

func TestParse_InvalidToken(t *testing.T) {
	m := newManager()
	_, err := m.Parse("not.a.token")
	assert.Error(t, err)
}

func TestParse_ExpiredToken(t *testing.T) {
	// Manager with 1ns expiry
	m := pkgjwt.NewManager(
		"test-secret-that-is-long-enough-ok",
		time.Nanosecond,
		time.Nanosecond,
	)
	token, err := m.SignAccess("u", "u@u.com", "user")
	require.NoError(t, err)

	time.Sleep(5 * time.Millisecond)

	_, err = m.Parse(token)
	assert.Error(t, err)
}
