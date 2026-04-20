package csv_pipeline

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Parse reads all data rows (skipping the header) from r.
// Returns an error only on structural CSV problems (wrong column count, IO errors).
// Per-field validation is NOT done here.
func Parse(r io.Reader) ([]RawRow, error) {
	reader := csv.NewReader(r)
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}
	if len(header) != 4 {
		return nil, fmt.Errorf("expected 4 columns, got %d", len(header))
	}

	var rows []RawRow
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("reading row: %w", err)
		}
		// TODO: map record fields to RawRow and append
		_ = record
	}
	return rows, nil
}
