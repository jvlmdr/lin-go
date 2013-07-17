package zvec

// Miscellaneous operations you can do with a Mutable vector.

import "fmt"

// Copies src into dst.
// Must have same size.
func Copy(dst Mutable, src Const) {
	if dst.Size() != src.Size() {
		panic(fmt.Sprintf("Vectors are different sizes (%d and %d)", dst.Size(), src.Size()))
	}
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, src.At(i))
	}
}

// Computes cumulative sum.
// src and dst must have same size.
//
// Result is such that
//	dst.At(0) == src.At(0)
//	dst.At(src.Size() - 1) == Sum(src)
func CumSum(dst Mutable, src Const) {
	if dst.Size() != src.Size() {
		panic(fmt.Sprintf("Vectors are different sizes (%d and %d)", dst.Size(), src.Size()))
	}

	var total complex128
	for i := 0; i < src.Size(); i++ {
		total += src.At(i)
		dst.Set(i, total)
	}
}
