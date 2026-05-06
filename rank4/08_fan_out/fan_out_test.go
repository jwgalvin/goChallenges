package fan_out

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestFetchAll_AllSucceed(t *testing.T) {
	urls := []string{"a", "b", "c", "d", "e"}
	fetch := func(_ context.Context, url string) (string, error) {
		return "body:" + url, nil
	}

	results := FetchAll(context.Background(), fetch, urls, 3)

	if len(results) != len(urls) {
		t.Fatalf("got %d results, want %d", len(results), len(urls))
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("url %s: unexpected error %v", r.URL, r.Err)
		}
		if r.Body != "body:"+r.URL {
			t.Errorf("url %s: body = %q, want %q", r.URL, r.Body, "body:"+r.URL)
		}
	}
}

func TestFetchAll_ErrorPropagation(t *testing.T) {
	boom := errors.New("boom")
	fetch := func(_ context.Context, url string) (string, error) {
		if url == "bad" {
			return "", boom
		}
		return "ok", nil
	}

	results := FetchAll(context.Background(), fetch, []string{"good", "bad"}, 2)

	if len(results) != 2 {
		t.Fatalf("got %d results, want 2", len(results))
	}
	var errCount int
	for _, r := range results {
		if r.Err != nil {
			errCount++
			if !errors.Is(r.Err, boom) {
				t.Errorf("url %s: err = %v, want boom", r.URL, r.Err)
			}
		}
	}
	if errCount != 1 {
		t.Errorf("got %d errors, want 1", errCount)
	}
}

func TestFetchAll_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var started atomic.Int32
	fetch := func(ctx context.Context, url string) (string, error) {
		started.Add(1)
		select {
		case <-time.After(10 * time.Second):
			return "ok", nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}

	urls := make([]string, 20)
	for i := range urls {
		urls[i] = fmt.Sprintf("url-%d", i)
	}

	done := make(chan []Result)
	go func() {
		done <- FetchAll(ctx, fetch, urls, 3)
	}()

	// wait until workers are running, then cancel
	for started.Load() < 3 {
		time.Sleep(time.Millisecond)
	}
	cancel()

	select {
	case results := <-done:
		// FetchAll must return (not hang) after cancellation
		if len(results) >= len(urls) {
			t.Errorf("expected fewer than %d results after cancellation, got %d", len(urls), len(results))
		}
	case <-time.After(2 * time.Second):
		t.Fatal("FetchAll did not return after context cancellation — goroutine leak suspected")
	}
}

func TestFetchAll_ConcurrencyCap(t *testing.T) {
	var inFlight atomic.Int32
	var peak atomic.Int32
	concurrency := 3

	fetch := func(_ context.Context, url string) (string, error) {
		cur := inFlight.Add(1)
		defer inFlight.Add(-1)
		for {
			p := peak.Load()
			if cur <= p || peak.CompareAndSwap(p, cur) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		return "ok", nil
	}

	urls := make([]string, 15)
	for i := range urls {
		urls[i] = fmt.Sprintf("url-%d", i)
	}

	FetchAll(context.Background(), fetch, urls, concurrency)

	if p := peak.Load(); p > int32(concurrency) {
		t.Errorf("peak concurrency = %d, want <= %d", p, concurrency)
	}
}

func TestFetchAll_EmptyInput(t *testing.T) {
	results := FetchAll(context.Background(), nil, []string{}, 5)
	if len(results) != 0 {
		t.Errorf("got %d results for empty input, want 0", len(results))
	}
}
