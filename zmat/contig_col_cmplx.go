package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A Contiguous) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A Contiguous) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A Contiguous) Imag() mat.Mutable { return MutableImag(A) }
