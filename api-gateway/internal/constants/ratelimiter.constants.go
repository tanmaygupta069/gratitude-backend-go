// middlewares/ratelimiter.go

package constants

import (
	"sync"
	"golang.org/x/time/rate"
)

// Limiter defines a rate limiter for each client.
type Limiter struct {
	mu      sync.Mutex
	clients map[string]*rate.Limiter
	rate    rate.Limit
	burst   int
}

// NewLimiter creates a new Limiter with the given rate and burst size.
func NewLimiter(r rate.Limit, b int) *Limiter {
	return &Limiter{
		clients: make(map[string]*rate.Limiter),
		rate:    r,
		burst:   b,
	}
}

// AddClient returns the limiter associated with a specific client.
func (l *Limiter) AddClient(clientIP string) *rate.Limiter {
	l.mu.Lock()
	defer l.mu.Unlock()

	limiter, exists := l.clients[clientIP]
	if !exists {
		limiter = rate.NewLimiter(l.rate, l.burst)
		l.clients[clientIP] = limiter
	}

	return limiter
}


