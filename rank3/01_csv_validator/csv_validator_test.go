package csv_validator

import (
	"strings"
	"testing"
)

func ptr(f float64) *float64 { return &f }

var standardRules = []FieldRule{
	{Name: "name", Required: true, Type: TypeString},
	{Name: "age", Required: true, Type: TypeInt, Min: ptr(0), Max: ptr(150)},
	{Name: "score", Required: false, Type: TypeFloat, Min: ptr(0), Max: ptr(100)},
}

func TestValidateCSV_AllValid(t *testing.T) {
	csv := "name,age,score\nAlice,30,95.5\nBob,25,88.0\n"
	errs, err := ValidateCSV(strings.NewReader(csv), standardRules)
	if err != nil {
		t.Fatalf("unexpected structural error: %v", err)
	}
	if len(errs) != 0 {
		t.Errorf("expected no validation errors, got: %v", errs)
	}
}

func TestValidateCSV_OptionalFieldMayBeEmpty(t *testing.T) {
	// score is optional — an empty score column is allowed
	csv := "name,age,score\nAlice,30,\n"
	errs, err := ValidateCSV(strings.NewReader(csv), standardRules)
	if err != nil {
		t.Fatalf("unexpected structural error: %v", err)
	}
	if len(errs) != 0 {
		t.Errorf("empty optional field should not produce an error, got: %v", errs)
	}
}

func TestValidateCSV_MissingRequiredField(t *testing.T) {
	// name is empty on line 2
	csv := "name,age,score\n,30,50.0\nBob,25,75.0\n"
	errs, err := ValidateCSV(strings.NewReader(csv), standardRules)
	if err != nil {
		t.Fatalf("unexpected structural error: %v", err)
	}
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, got %d: %v", len(errs), errs)
	}
	if errs[0].Line != 2 || errs[0].Field != "name" {
		t.Errorf("error = %+v; want line 2, field \"name\"", errs[0])
	}
}

func TestValidateCSV_WrongType(t *testing.T) {
	// age is "thirty" instead of an int
	csv := "name,age,score\nAlice,thirty,80.0\n"
	errs, _ := ValidateCSV(strings.NewReader(csv), standardRules)
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, got %d: %v", len(errs), errs)
	}
	if errs[0].Field != "age" {
		t.Errorf("expected error on field \"age\", got %q", errs[0].Field)
	}
}

func TestValidateCSV_OutOfRange(t *testing.T) {
	// age 200 exceeds max 150; score 110 exceeds max 100
	csv := "name,age,score\nAlice,200,110.0\n"
	errs, _ := ValidateCSV(strings.NewReader(csv), standardRules)
	if len(errs) != 2 {
		t.Fatalf("expected 2 errors, got %d: %v", len(errs), errs)
	}
	fields := map[string]bool{}
	for _, e := range errs {
		fields[e.Field] = true
	}
	if !fields["age"] || !fields["score"] {
		t.Errorf("expected errors on 'age' and 'score', got fields: %v", fields)
	}
}

func TestValidateCSV_MultipleErrorsAcrossLines(t *testing.T) {
	// line 2: name missing, age invalid type
	// line 3: score out of range
	// line 4: valid
	csv := "name,age,score\n,oops,50.0\nBob,40,200.0\nCarol,35,90.0\n"
	errs, _ := ValidateCSV(strings.NewReader(csv), standardRules)
	if len(errs) != 3 {
		t.Errorf("expected 3 errors, got %d: %v", len(errs), errs)
	}
}

func TestValidateCSV_EmptyInput(t *testing.T) {
	_, err := ValidateCSV(strings.NewReader(""), standardRules)
	if err == nil {
		t.Error("expected structural error for empty input (no header row)")
	}
}

func TestValidateCSV_HeaderMissingColumn(t *testing.T) {
	// CSV only has "name" but rules also require "age"
	csv := "name\nAlice\n"
	_, err := ValidateCSV(strings.NewReader(csv), standardRules)
	if err == nil {
		t.Error("expected structural error when header is missing a required column")
	}
}

func TestValidateCSV_ColumnsInAnyOrder(t *testing.T) {
	// Columns may appear in a different order than the rules slice
	csv := "score,name,age\n85.0,Alice,30\n"
	errs, err := ValidateCSV(strings.NewReader(csv), standardRules)
	if err != nil {
		t.Fatalf("unexpected structural error: %v", err)
	}
	if len(errs) != 0 {
		t.Errorf("expected no errors for valid reordered CSV, got: %v", errs)
	}
}
