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

// FetchAll fetches all urls concurrently with at most `concurrency` in-flight
// at once. It returns one Result per URL in any order. If ctx is canceled,
// in-flight work is abandoned and no new work is started; the function still
// returns (with fewer results than urls).
func FetchAll(ctx context.Context, fetch FetchFunc, urls []string, concurrency int) []Result {
	// TODO: implement
	return nil
}
