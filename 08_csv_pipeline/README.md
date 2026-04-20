# Task: CSV Import Pipeline

Build a product-import pipeline that reads a CSV, validates each row,
transforms it into a domain struct, and collects per-row errors instead
of stopping at the first failure.

## CSV format

```
sku,name,price_cents,stock
ABC-1,Widget,999,50
BAD,,0,-1
```

## Requirements

1. `parser.go`    — parse raw CSV rows into `RawRow` (no validation here)
2. `validator.go` — validate a `RawRow`, return a `[]ValidationError` (one per field violation)
3. `transformer.go` — convert a valid `RawRow` into a `Product`
4. `pipeline.go`  — wire them together; return `Result{Products, Errors}` where `Errors` are
                    `RowError{Line int, Errs []ValidationError}`

## Validation rules
- `sku`         non-empty, matches `[A-Z]+-[0-9]+`
- `name`        non-empty
- `price_cents` integer > 0
- `stock`       integer >= 0

## Signals being tested
- Composable, single-responsibility functions
- Collecting multiple errors per row (not early-exit)
- Clean error types (not fmt.Errorf strings)
- Table-driven tests across happy path, partial failure, full failure
