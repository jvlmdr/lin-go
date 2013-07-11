package vec

// Miscellaneous operations you can do with a Mutable vector.

func Copy(dst Mutable, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, src.At(i))
	}
}
