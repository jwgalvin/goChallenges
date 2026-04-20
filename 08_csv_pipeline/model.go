package csv_pipeline

// RawRow holds the unvalidated string values from one CSV line.
type RawRow struct {
	SKU        string
	Name       string
	PriceCents string
	Stock      string
}

// Product is the validated, typed result.
type Product struct {
	SKU        string
	Name       string
	PriceCents int
	Stock      int
}

// ValidationError describes a single field-level problem.
type ValidationError struct {
	Field   string
	Message string
}

// RowError bundles a 1-based line number with its validation errors.
type RowError struct {
	Line int
	Errs []ValidationError
}

// Result is the output of the pipeline.
type Result struct {
	Products []Product
	Errors   []RowError
}
