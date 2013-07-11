package mat

// Constructs default matrix.
func Make(rows, cols int) Mutable {
	return MakeContiguous(rows, cols)
}

// Constructs default matrix.
func MakeCopy(B Const) Mutable {
	return MakeContiguousCopy(B)
}
