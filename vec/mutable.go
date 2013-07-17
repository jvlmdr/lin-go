package vec

// Miscellaneous operations you can do with a Mutable vector.

import "fmt"

// Copies src into dst.
// Must have same size.
func Copy(dst Mutable, src Const) {
	if dst.Len() != src.Len() {
		panic(fmt.Sprintf("Vectors are different sizes (%d and %d)", dst.Len(), src.Len()))
	}
	for i := 0; i < src.Len(); i++ {
		dst.Set(i, src.At(i))
	}
}

// Computes cumulative sum.
// src and dst must have same size.
// Can be done in-place.
//
// Result is such that
//	dst.At(0) == src.At(0)
//	dst.At(src.Len() - 1) == Sum(src)
func CumSum(dst Mutable, src Const) {
	if dst.Len() != src.Len() {
		panic(fmt.Sprintf("Vectors are different sizes (%d and %d)", dst.Len(), src.Len()))
	}

	var total float64
	for i := 0; i < src.Len(); i++ {
		total += src.At(i)
		dst.Set(i, total)
	}
}
