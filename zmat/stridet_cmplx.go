package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A StrideT) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A StrideT) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A StrideT) Imag() mat.Mutable { return MutableImag(A) }
