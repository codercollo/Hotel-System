// Package crypto provides HMAC-SHA256 signing utilities used primarily
// to verify webhook payloads from external payment providers.
package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign returns the HMAC-SHA256 hex-encoded signature of payload using secret.
func Sign(secret, payload string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

// Verify reports whether the provided signature is the valid HMAC-SHA256
// signature of payload using secret. Comparison is constant-time.
func Verify(secret, payload, signature string) bool {
	expected := Sign(secret, payload)
	return hmac.Equal([]byte(expected), []byte(signature))
}
