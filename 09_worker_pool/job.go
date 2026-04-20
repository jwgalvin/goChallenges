package worker_pool

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
