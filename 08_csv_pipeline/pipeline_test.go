package csv_pipeline

import (
	"strings"
	"testing"
)

const header = "sku,name,price_cents,stock\n"

func TestRun_AllValid(t *testing.T) {
	input := header + "ABC-1,Widget,999,50\nXYZ-99,Gadget,1500,0\n"
	result, err := Run(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Products) != 2 {
		t.Errorf("products = %d, want 2", len(result.Products))
	}
	if len(result.Errors) != 0 {
		t.Errorf("errors = %v, want none", result.Errors)
	}
	if result.Products[0].PriceCents != 999 {
		t.Errorf("PriceCents = %d, want 999", result.Products[0].PriceCents)
	}
}

func TestRun_PartialFailure(t *testing.T) {
	// row 1 valid, row 2 invalid (empty name, zero price)
	input := header + "ABC-1,Widget,999,50\n,,-1,0\n"
	result, err := Run(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Products) != 1 {
		t.Errorf("products = %d, want 1", len(result.Products))
	}
	if len(result.Errors) != 1 {
		t.Fatalf("errors = %d, want 1", len(result.Errors))
	}
	if result.Errors[0].Line != 2 {
		t.Errorf("error line = %d, want 2", result.Errors[0].Line)
	}
	// should report multiple field errors on the same row
	if len(result.Errors[0].Errs) < 2 {
		t.Errorf("expected >= 2 field errors, got %d", len(result.Errors[0].Errs))
	}
}

func TestRun_AllFail(t *testing.T) {
	input := header + "bad sku,,-1,-5\n"
	result, err := Run(strings.NewReader(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Products) != 0 {
		t.Errorf("products = %d, want 0", len(result.Products))
	}
	if len(result.Errors) != 1 {
		t.Errorf("errors = %d, want 1", len(result.Errors))
	}
}

func TestValidate_SKUPattern(t *testing.T) {
	cases := []struct {
		sku     string
		wantErr bool
	}{
		{"ABC-1", false},
		{"ABC-123", false},
		{"abc-1", true},  // lowercase
		{"ABC1", true},   // missing dash
		{"", true},
		{"ABC-", true},
	}
	for _, c := range cases {
		row := RawRow{SKU: c.sku, Name: "Widget", PriceCents: "100", Stock: "0"}
		errs := Validate(row)
		hasErr := false
		for _, e := range errs {
			if e.Field == "sku" {
				hasErr = true
			}
		}
		if hasErr != c.wantErr {
			t.Errorf("SKU %q: wantErr=%v, got errors=%v", c.sku, c.wantErr, errs)
		}
	}
}
