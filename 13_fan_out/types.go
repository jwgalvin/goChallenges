package fan_out

import "context"

// FetchFunc is the function used to fetch a single URL.
// Injected so tests don't make real HTTP calls.
type FetchFunc func(ctx context.Context, url string) (string, error)

// Result holds the outcome of fetching one URL.
type Result struct {
	URL  string
	Body string
	Err  error
}
