package rate_limiter

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(3, 100*time.Millisecond)

	// first 3 should be allowed
	for i := range 3 {
		if !rl.Allow() {
			t.Fatalf("request %d should be allowed", i+1)
		}
	}
	// 4th should be denied
	if rl.Allow() {
		t.Fatal("4th request should be denied")
	}

	// after window expires, counter resets
	time.Sleep(110 * time.Millisecond)
	if !rl.Allow() {
		t.Fatal("request after window reset should be allowed")
	}
}
