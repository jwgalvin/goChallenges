package worker_pool

import "errors"

var ErrPoolStopped = errors.New("pool is stopped")

// Job is a unit of work.
type Job struct {
	ID      string
	Payload func() (any, error)
}

// Result holds the outcome of a single Job.
type Result struct {
	JobID string
	Value any
	Err   error
}

// Pool manages N worker goroutines.
type Pool struct {
	// TODO: add fields (jobs channel, results channel, wg, stopped flag, ...)
}

// New creates a Pool with the given number of workers and starts them.
// Results are available on the channel returned by Results().
func New(workers int) *Pool {
	// TODO: implement
	return &Pool{}
}

// Results returns a receive-only channel that emits one Result per completed Job.
func (p *Pool) Results() <-chan Result {
	// TODO: implement
	return nil
}

// Submit enqueues a job. Blocks if the queue is full.
// Returns ErrPoolStopped if Stop has been called.
func (p *Pool) Submit(job Job) error {
	// TODO: implement
	return ErrPoolStopped
}

// Stop signals workers to finish, waits for all in-flight jobs to complete,
// then closes the results channel.
func (p *Pool) Stop() {
	// TODO: implement
}
