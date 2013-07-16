package zvec

// Makes a default mutable vector.
func Make(n int) Slice {
	return MakeSlice(n)
}

// Makes a default mutable vector and copies x into it.
func MakeCopy(x Const) Slice {
	return MakeSliceCopy(x)
}
