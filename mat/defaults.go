package mat

// Constructs default matrix.
func Make(rows, cols int) Contig {
	return MakeContig(rows, cols)
}

// Constructs default matrix.
func MakeCopy(B Const) Contig {
	return MakeContigCopy(B)
}
