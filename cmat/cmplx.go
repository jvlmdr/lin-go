package cmat

import (
	"github.com/jackvalmadre/lin-go/mat"
	"math/cmplx"
)

// Creates a complex matrix from a real matrix.
func NewReal(a mat.Const) *Mat {
	m, n := a.Dims()
	b := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(i, j, complex(a.At(i, j), 0))
		}
	}
	return b
}

// Creates a complex matrix from a real matrix.
//
// Panics if a and b are not the same size.
func NewRealImag(a, b mat.Const) *Mat {
	if err := errIfDimsNotEq(a, b); err != nil {
		panic(err)
	}

	m, n := a.Dims()
	c := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c.Set(i, j, complex(a.At(i, j), b.At(i, j)))
		}
	}
	return c
}

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

// Extracts the real part of a complex matrix.
func Real(a Const) *mat.Mat {
	m, n := a.Dims()
	b := mat.New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(i, j, real(a.At(i, j)))
		}
	}
	return b
}

// Extracts the imaginary part of a complex matrix.
func Imag(a Const) *mat.Mat {
	m, n := a.Dims()
	b := mat.New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(i, j, imag(a.At(i, j)))
		}
	}
	return b
}
