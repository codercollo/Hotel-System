// Package idgen generates sortable unique identifiers.
// ULIDs are used as primary keys throughout the platform for sortability
// and URL-friendliness. UUIDs are available for external integration contexts.
package idgen

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// NewULID returns a new ULID string (26 chars, sortable, URL-safe).
func NewULID() string {
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

// NewUUID returns a new UUID v4 string.
func NewUUID() string {
	return uuid.New().String()
}
