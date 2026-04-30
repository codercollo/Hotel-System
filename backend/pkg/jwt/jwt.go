// Package jwt provides sign, parse, and refresh helpers for HMAC-SHA256 JWTs.
// It wraps github.com/golang-jwt/jwt/v5 with typed claims and error mapping.
package jwt

import (
	"errors"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/golang-jwt/jwt/v5"
)

// TokenType distinguishes access from refresh tokens in the claims.
type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

// Claims is the custom JWT payload stored in every token.
type Claims struct {
	UserID    string    `json:"uid"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

// Manager handles signing and parsing JWTs with a shared secret.
type Manager struct {
	secret        []byte
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

// NewManager creates a JWT manager from raw configuration values.
func NewManager(secret string, accessExpiry, refreshExpiry time.Duration) *Manager {
	return &Manager{
		secret:        []byte(secret),
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

// SignAccess creates a signed access token for the given user.
func (m *Manager) SignAccess(userID, email, role string) (string, error) {
	return m.sign(userID, email, role, TokenTypeAccess, m.accessExpiry)
}

// SignRefresh creates a signed refresh token.
func (m *Manager) SignRefresh(userID, email, role string) (string, error) {
	return m.sign(userID, email, role, TokenTypeRefresh, m.refreshExpiry)
}

// Parse validates a token string and returns the embedded claims.
// Returns typed apierror values for expired / invalid tokens.
func (m *Manager) Parse(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apierror.ErrTokenInvalid
			}
			return m.secret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, apierror.ErrTokenExpired
		}
		return nil, apierror.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, apierror.ErrTokenInvalid
	}

	return claims, nil
}

// ParseRefresh validates that the token is a valid refresh token.
func (m *Manager) ParseRefresh(tokenStr string) (*Claims, error) {
	claims, err := m.Parse(tokenStr)
	if err != nil {
		return nil, err
	}
	if claims.TokenType != TokenTypeRefresh {
		return nil, apierror.ErrTokenInvalid
	}
	return claims, nil
}

func (m *Manager) sign(userID, email, role string, tt TokenType, expiry time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:    userID,
		Email:     email,
		Role:      role,
		TokenType: tt,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(m.secret)
	if err != nil {
		return "", apierror.Internal(err)
	}
	return signed, nil
}
