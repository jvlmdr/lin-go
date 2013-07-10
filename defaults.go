package mat

// Default (column-major contiguous) matrix.
func Make(rows, cols int) Mutable {
	return MakeContiguous(rows, cols)
}
