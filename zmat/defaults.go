package zmat

// Constructs default matrix.
func Make(rows, cols int) Contiguous {
	return MakeContiguous(rows, cols)
}

// Constructs default matrix.
func MakeCopy(B Const) Contiguous {
	return MakeContiguousCopy(B)
}
