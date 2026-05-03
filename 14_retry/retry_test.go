package retry

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry_SucceedsOnFirstAttempt(t *testing.T) {
	calls := 0
	err := Retry(context.Background(), 3, func() error {
		calls++
		return nil
	})
	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
	if calls != 1 {
		t.Errorf("fn called %d times, want 1", calls)
	}
}

func TestRetry_SucceedsAfterFailures(t *testing.T) {
	calls := 0
	err := Retry(context.Background(), 5, func() error {
		calls++
		if calls < 3 {
			return fmt.Errorf("not ready yet")
		}
		return nil
	})
	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
	if calls != 3 {
		t.Errorf("fn called %d times, want 3", calls)
	}
}

func TestRetry_ExhaustsAttempts(t *testing.T) {
	sentinel := errors.New("always fails")
	calls := 0
	err := Retry(context.Background(), 3, func() error {
		calls++
		return sentinel
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, sentinel) {
		t.Errorf("err = %v, want to wrap sentinel", err)
	}
	if calls != 3 {
		t.Errorf("fn called %d times, want 3", calls)
	}
}

func TestRetry_ErrorWrapsAttemptNumber(t *testing.T) {
	err := Retry(context.Background(), 2, func() error {
		return errors.New("bad")
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	// The final error should mention the attempt number somewhere
	for _, want := range []string{"2", "attempt"} {
		if !containsStr(msg, want) {
			t.Logf("err message: %q", msg)
			t.Errorf("expected error message to contain %q", want)
		}
	}
}

func TestRetry_ContextCanceledDuringWait(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	calls := 0
	start := time.Now()

	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	err := Retry(ctx, 10, func() error {
		calls++
		return errors.New("fail")
	})

	elapsed := time.Since(start)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("err = %v, want context.Canceled", err)
	}
	// Should have returned well before attempting all 10 retries
	if elapsed > 5*time.Second {
		t.Errorf("took %v, should have returned shortly after cancellation", elapsed)
	}
	if calls >= 10 {
		t.Errorf("fn called %d times — cancellation didn't stop retries", calls)
	}
}

func TestRetry_ZeroAttemptsReturnsNil(t *testing.T) {
	called := false
	err := Retry(context.Background(), 0, func() error {
		called = true
		return errors.New("should not be called")
	})
	if err != nil {
		t.Errorf("err = %v, want nil for 0 attempts", err)
	}
	if called {
		t.Error("fn should not be called when maxAttempts is 0")
	}
}

func containsStr(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 ||
		func() bool {
			for i := 0; i <= len(s)-len(sub); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		}())
}
