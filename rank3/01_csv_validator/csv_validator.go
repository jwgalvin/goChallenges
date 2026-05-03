package csv_validator

import (
	"fmt"
	"io"
)

type FieldType string

const (
	TypeString FieldType = "string"
	TypeInt    FieldType = "int"
	TypeFloat  FieldType = "float"
)

// FieldRule describes validation rules for one CSV column.
type FieldRule struct {
	Name     string
	Required bool
	Type     FieldType
	Min, Max *float64 // nil means no bound; only applies to int and float types
}

// ValidationError is a single field-level validation failure.
type ValidationError struct {
	Line  int
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("line %d, field %q: %s", e.Line, e.Field, e.Msg)
}

// ValidateCSV reads CSV data from r and validates each data row against rules.
// The first row is treated as a header; data rows start at line 2.
//
// It returns a non-nil error only for structural problems: unreadable input,
// or the header missing a column named in rules. Field-level failures are
// returned in the []ValidationError slice and do not stop parsing.
//
// Hint: use encoding/csv to read rows; strconv to parse numbers.
func ValidateCSV(r io.Reader, rules []FieldRule) ([]ValidationError, error) {
	// TODO: implement
	return nil, nil
}
