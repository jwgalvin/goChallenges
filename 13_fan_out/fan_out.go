package fan_out

import "context"

// FetchAll fetches all urls concurrently with at most `concurrency` in-flight
// at once. It returns one Result per URL in any order. If ctx is canceled,
// in-flight work is abandoned and no new work is started; the function still
// returns (with fewer results than urls).
func FetchAll(ctx context.Context, fetch FetchFunc, urls []string, concurrency int) []Result {
	// TODO: implement
	return nil
}
