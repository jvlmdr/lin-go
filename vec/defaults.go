package vec

// Makes a default mutable vector.
func Make(n int) Mutable {
	return MakeSlice(n)
}

// Makes a default mutable vector and copies x into it.
func MakeCopy(x Const) Mutable {
	return MakeSliceCopy(x)
}
