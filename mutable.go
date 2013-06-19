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

func AddTo(dst Mutable, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, dst.At(i)+src.At(i))
	}
}

func SubtractFrom(dst Mutable, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, dst.At(i)-src.At(i))
	}
}

func ScaleAndCopyTo(dst Mutable, a float64, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, a*src.At(i))
	}
}
