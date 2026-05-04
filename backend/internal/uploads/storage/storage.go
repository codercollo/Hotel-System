// Package storage defines the file storage abstraction.
// Implementations: local.go (dev), s3.go (production).
package storage

import (
	"context"
	"io"
)

// UploadResult is returned after a successful upload.
type UploadResult struct {
	Key  string // storage key / path
	URL  string // public URL (if applicable)
	Size int64
}

// Storage is the interface every backend must satisfy.
type Storage interface {
	// Upload stores a file and returns a result with its key and URL.
	Upload(ctx context.Context, key string, r io.Reader, size int64, mimeType string) (*UploadResult, error)

	// Delete removes a file by key.
	Delete(ctx context.Context, key string) error

	// URL returns the public URL for an existing key.
	URL(key string) string
}
