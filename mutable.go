package vec

func Fill(x Mutable, v float64) {
	for i := 0; i < x.Size(); i++ {
		x.Set(i, v)
	}
}

func CopyTo(dst Mutable, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, src.At(i))
	}
}
