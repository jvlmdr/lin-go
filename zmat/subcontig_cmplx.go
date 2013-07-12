package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A SubContiguous) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A SubContiguous) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A SubContiguous) Imag() mat.Mutable { return MutableImag(A) }
