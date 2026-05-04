package storage

import (
	"context"
	"fmt"
	"io"
)

// S3Storage stores files in AWS S3 (or any S3-compatible service).
// Full implementation in Phase 7 using the AWS SDK.
type S3Storage struct {
	bucket   string
	region   string
	endpoint string // empty = AWS; set for Cloudflare R2, MinIO, etc.
	baseURL  string
}

// NewS3Storage constructs an S3 storage backend stub.
func NewS3Storage(bucket, region, endpoint, baseURL string) *S3Storage {
	return &S3Storage{bucket: bucket, region: region, endpoint: endpoint, baseURL: baseURL}
}

// Upload puts a file to S3. Phase 7 wires the real AWS SDK call.
func (s *S3Storage) Upload(_ context.Context, key string, _ io.Reader, _ int64, _ string) (*UploadResult, error) {
	return nil, fmt.Errorf("s3: not yet implemented — Phase 7")
}

func (s *S3Storage) Delete(_ context.Context, key string) error {
	return fmt.Errorf("s3: not yet implemented — Phase 7")
}

func (s *S3Storage) URL(key string) string {
	if s.baseURL != "" {
		return fmt.Sprintf("%s/%s", s.baseURL, key)
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
}
