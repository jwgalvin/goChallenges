package csv_pipeline

// Validate checks every field in row and returns all violations found.
// An empty slice means the row is valid.
func Validate(row RawRow) []ValidationError {
	// TODO: implement all four rules (see README)
	// Hint: use regexp.MustCompile at package level for the SKU pattern
	return nil
}
