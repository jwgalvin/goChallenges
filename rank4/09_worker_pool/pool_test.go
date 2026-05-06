package worker_pool

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func collectResults(p *Pool, n int) []Result {
	var results []Result
	for r := range p.Results() {
		results = append(results, r)
		if len(results) == n {
			break
		}
	}
	return results
}

func TestPool_BasicExecution(t *testing.T) {
	p := New(3)

	for i := range 5 {
		id := fmt.Sprintf("job-%d", i)
		val := i
		if err := p.Submit(Job{ID: id, Payload: func() (any, error) {
			return val * 2, nil
		}}); err != nil {
			t.Fatalf("Submit: %v", err)
		}
	}

	results := collectResults(p, 5)
	p.Stop()

	if len(results) != 5 {
		t.Errorf("got %d results, want 5", len(results))
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("job %s errored: %v", r.JobID, r.Err)
		}
	}
}

func TestPool_ErrorPropagation(t *testing.T) {
	p := New(2)
	boom := errors.New("boom")
	_ = p.Submit(Job{ID: "fail", Payload: func() (any, error) {
		return nil, boom
	}})

	results := collectResults(p, 1)
	p.Stop()

	if !errors.Is(results[0].Err, boom) {
		t.Errorf("err = %v, want %v", results[0].Err, boom)
	}
}

func TestPool_PanicRecovery(t *testing.T) {
	p := New(2)
	_ = p.Submit(Job{ID: "panic-job", Payload: func() (any, error) {
		panic("something went wrong")
	}})

	results := collectResults(p, 1)
	p.Stop()

	if results[0].Err == nil {
		t.Error("expected panic to be converted to an error")
	}
}

func TestPool_StopDrainsInFlight(t *testing.T) {
	p := New(2)
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := range n {
		id := fmt.Sprintf("j%d", i)
		_ = p.Submit(Job{ID: id, Payload: func() (any, error) {
			time.Sleep(5 * time.Millisecond)
			return nil, nil
		}})
	}

	go func() {
		for range p.Results() {
			wg.Done()
		}
	}()

	p.Stop() // must not return until all 10 are done
	wg.Wait()
}

func TestPool_SubmitAfterStop(t *testing.T) {
	p := New(1)
	p.Stop()
	if err := p.Submit(Job{ID: "late"}); !errors.Is(err, ErrPoolStopped) {
		t.Errorf("err = %v, want ErrPoolStopped", err)
	}
}
