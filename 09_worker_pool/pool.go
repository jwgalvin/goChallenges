package worker_pool

import (
	"errors"
	"fmt"
	"sync"
)

var ErrPoolStopped = errors.New("pool is stopped")

// Pool manages N worker goroutines.
type Pool struct {
	jobs    chan Job
	results chan Result
	wg      sync.WaitGroup
	workers sync.WaitGroup
	mu      sync.Mutex
	stopped bool
	once    sync.Once
}

// New creates a Pool with the given number of workers and starts them.
// Results are available on the channel returned by Results().
func New(workers int) *Pool {
	p := &Pool{
		jobs:    make(chan Job, workers*2),
		results: make(chan Result, workers*2),
	}
	p.workers.Add(workers)
	for i := 0; i < workers; i++ {
		go p.worker()
	}
	return p
}

func (p *Pool) worker() {
	defer p.workers.Done()
	for job := range p.jobs {
		result := p.runJob(job)
		p.results <- result
		p.wg.Done()
	}
}

func (p *Pool) runJob(job Job) (r Result) {
	r.JobID = job.ID
	defer func() {
		if rec := recover(); rec != nil {
			r.Err = fmt.Errorf("panic: %v", rec)
		}
	}()
	r.Value, r.Err = job.Payload()
	return
}

// Results returns a receive-only channel that emits one Result per completed Job.
func (p *Pool) Results() <-chan Result {
	return p.results
}

// Submit enqueues a job. Blocks if the queue is full.
// Returns ErrPoolStopped if Stop has been called.
func (p *Pool) Submit(job Job) error {
	p.mu.Lock()
	if p.stopped {
		p.mu.Unlock()
		return ErrPoolStopped
	}
	p.wg.Add(1)
	p.mu.Unlock()

	p.jobs <- job
	return nil
}

// Stop signals workers to finish, waits for all in-flight jobs to complete,
// then closes the results channel.
func (p *Pool) Stop() {
	p.once.Do(func() {
		p.mu.Lock()
		p.stopped = true
		p.mu.Unlock()

		p.wg.Wait()
		close(p.jobs)
		p.workers.Wait()
		close(p.results)
	})
}
