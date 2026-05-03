package retry

import "context"

// Retry calls fn up to maxAttempts times. It waits between attempts using
// exponential backoff starting at baseDelay, doubling each time.
// If ctx is canceled during a wait, it returns immediately with ctx.Err().
// Each error is wrapped with the attempt number using fmt.Errorf.
// Returns nil as soon as fn returns nil.
func Retry(ctx context.Context, maxAttempts int, fn func() error) error {
	// TODO: implement
	return nil
}
