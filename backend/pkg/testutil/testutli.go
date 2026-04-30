// Package testutil provides shared helpers for unit and integration tests.
package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/cache"
)

// MustMarshal encodes v to JSON and fails the test on error.
func MustMarshal(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("testutil.MustMarshal: %v", err)
	}
	return b
}

// NewJSONRequest creates an HTTP request with a JSON body.
func NewJSONRequest(t *testing.T, method, path string, body any) *http.Request {
	t.Helper()
	var r io.Reader
	if body != nil {
		r = bytes.NewReader(MustMarshal(t, body))
	}
	req, err := http.NewRequest(method, path, r)
	if err != nil {
		t.Fatalf("testutil.NewJSONRequest: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

// NewRecorder wraps httptest.NewRecorder with a convenience decode method.
type Recorder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *Recorder {
	return &Recorder{httptest.NewRecorder()}
}

// DecodeBody decodes the response body JSON into dst.
func (r *Recorder) DecodeBody(t *testing.T, dst any) {
	t.Helper()
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		t.Fatalf("testutil.Recorder.DecodeBody: %v", err)
	}
}

// ─── In-Memory Cache ──────────────────────────────────────────────────────────

// MemCache is a thread-unsafe in-memory cache for use in tests.
type MemCache struct {
	data map[string]entry
}

type entry struct {
	value   string
	expires time.Time
}

func NewMemCache() *MemCache {
	return &MemCache{data: make(map[string]entry)}
}

func (m *MemCache) Get(_ context.Context, key string) (string, error) {
	e, ok := m.data[key]
	if !ok || (!e.expires.IsZero() && time.Now().After(e.expires)) {
		return "", nil
	}
	return e.value, nil
}

func (m *MemCache) Set(_ context.Context, key string, value any, ttl time.Duration) error {
	v := ""
	switch t := value.(type) {
	case string:
		v = t
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		v = string(b)
	}
	var exp time.Time
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}
	m.data[key] = entry{value: v, expires: exp}
	return nil
}

func (m *MemCache) Delete(_ context.Context, key string) error {
	delete(m.data, key)
	return nil
}

func (m *MemCache) Exists(_ context.Context, key string) (bool, error) {
	e, ok := m.data[key]
	if !ok {
		return false, nil
	}
	if !e.expires.IsZero() && time.Now().After(e.expires) {
		return false, nil
	}
	return true, nil
}

// Compile-time check: MemCache satisfies cache.Cache.
var _ cache.Cache = (*MemCache)(nil)
