package cmat

import "math/cmplx"

// Creates a conjugate-transposed copy of the matrix.
func H(a Const) *Mat {
	m, n := a.Dims()
	b := New(n, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(j, i, cmplx.Conj(a.At(i, j)))
		}
	}
	return b
}

// Creates a conjugated copy of the matrix.
func Conj(a Const) *Mat {
	m, n := a.Dims()
	b := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(i, j, cmplx.Conj(a.At(i, j)))
		}
	}
	return b
}
