# Task: Worker Pool

Build a bounded concurrent worker pool that processes jobs and collects results.

## Requirements

1. `pool.go`  — `Pool` runs exactly N goroutines; accepts `Job` values via `Submit`,
               returns `Result` values via a results channel; shuts down cleanly on `Stop`.
2. `job.go`   — define `Job` (an ID + a `func() (any, error)` payload) and `Result`
               (the Job ID, the return value, and any error).
3. `pool_test.go` — all tests pass with a race-free implementation (`go test -race`).

## Constraints
- `Submit` must block if the job queue is full (bounded channel, capacity = 2×workers).
- `Stop` must wait for all in-flight jobs to finish before returning.
- Workers must never panic — a job that panics should produce a Result with the recovered error.
- After `Stop`, calling `Submit` returns `ErrPoolStopped`.

## Signals being tested
- Goroutine lifecycle management (no leaks)
- Channel direction types in function signatures
- sync.WaitGroup usage
- Panic recovery
- Race-free state (test with -race flag)
