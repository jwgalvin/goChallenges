package csv_pipeline

// Transform converts a validated RawRow into a Product.
// It panics if called with an invalid row — callers must validate first.
func Transform(row RawRow) Product {
	// TODO: parse PriceCents and Stock to int, return Product
	return Product{}
}
