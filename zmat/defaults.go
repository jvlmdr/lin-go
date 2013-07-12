package zmat

// Constructs default matrix.
func Make(rows, cols int) ContiguousColMajor {
	return MakeContiguous(rows, cols)
}

// Constructs default matrix.
func MakeCopy(B Const) ContiguousColMajor {
	return MakeContiguousCopy(B)
}
