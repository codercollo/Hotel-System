package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LocalStorage stores files on the local filesystem. Development use only.
type LocalStorage struct {
	basePath string
	baseURL  string
}

// NewLocalStorage creates a local filesystem storage backend.
func NewLocalStorage(basePath, baseURL string) *LocalStorage {
	_ = os.MkdirAll(basePath, 0755)
	return &LocalStorage{basePath: basePath, baseURL: baseURL}
}

func (s *LocalStorage) Upload(_ context.Context, key string, r io.Reader, _ int64, _ string) (*UploadResult, error) {
	dest := filepath.Join(s.basePath, key)
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return nil, fmt.Errorf("local storage: mkdir: %w", err)
	}

	f, err := os.Create(dest)
	if err != nil {
		return nil, fmt.Errorf("local storage: create: %w", err)
	}
	defer f.Close()

	size, err := io.Copy(f, r)
	if err != nil {
		return nil, fmt.Errorf("local storage: write: %w", err)
	}

	return &UploadResult{Key: key, URL: s.URL(key), Size: size}, nil
}

func (s *LocalStorage) Delete(_ context.Context, key string) error {
	return os.Remove(filepath.Join(s.basePath, key))
}

func (s *LocalStorage) URL(key string) string {
	return fmt.Sprintf("%s/%s", s.baseURL, key)
}
