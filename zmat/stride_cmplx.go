package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A Stride) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A Stride) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A Stride) Imag() mat.Mutable { return MutableImag(A) }
