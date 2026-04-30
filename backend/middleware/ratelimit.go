package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/codercollo/hotel-system/backend/pkg/apierror"
	"github.com/codercollo/hotel-system/backend/pkg/response"
)

// RateLimiter provides a simple in-memory token-bucket rate limiter per IP.
// For production, swap this for a Redis-backed implementation.
type RateLimiter struct {
	mu       sync.Mutex
	buckets  map[string]*bucket
	requests int
	window   time.Duration
}

type bucket struct {
	count   int
	resetAt time.Time
}

// NewRateLimiter creates a limiter allowing `requests` requests per `window`.
func NewRateLimiter(requests int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		buckets:  make(map[string]*bucket),
		requests: requests,
		window:   window,
	}

	// Periodically clean up expired buckets.
	go func() {
		tick := time.NewTicker(window)
		for range tick.C {
			rl.cleanup()
		}
	}()

	return rl
}

// Middleware returns an http.Handler middleware enforcing the rate limit.
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		if !rl.allow(ip) {
			response.Error(w, apierror.TooManyRequests())
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, ok := rl.buckets[ip]
	if !ok || now.After(b.resetAt) {
		rl.buckets[ip] = &bucket{count: 1, resetAt: now.Add(rl.window)}
		return true
	}

	if b.count >= rl.requests {
		return false
	}

	b.count++
	return true
}

func (rl *RateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	for ip, b := range rl.buckets {
		if now.After(b.resetAt) {
			delete(rl.buckets, ip)
		}
	}
}
