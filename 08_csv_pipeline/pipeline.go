package csv_pipeline

import "io"

// Run executes the full import pipeline on r.
// Rows that pass validation become Products.
// Rows with validation errors become RowErrors (line numbers are 1-based, after header).
// A structural CSV error (bad format) is returned as a hard error.
func Run(r io.Reader) (Result, error) {
	// TODO: Parse → Validate → Transform, collecting per-row errors
	return Result{}, nil
}
