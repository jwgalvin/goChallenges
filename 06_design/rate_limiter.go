package design_practice

import "time"

// RateLimiter allows at most N requests per window duration.
// This is a fixed-window counter implementation.
type RateLimiter struct {
	// TODO: add fields
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	// TODO: implement
	return &RateLimiter{}
}

// Allow returns true if the request should be permitted.
func (r *RateLimiter) Allow() bool {
	// TODO: implement
	return false
}
