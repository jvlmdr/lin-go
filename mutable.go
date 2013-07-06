package vec

import "math/rand"

func Randn(x Mutable) {
	for i := 0; i < x.Size(); i++ {
		x.Set(i, rand.NormFloat64())
	}
}

func CopyTo(dst Mutable, src Const) {
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, src.At(i))
	}
}

func AddTo(x Mutable, y Const) {
	CopyTo(x, Plus(x, y))
}
