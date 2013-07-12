package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A SubContiguousRowMajor) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A SubContiguousRowMajor) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A SubContiguousRowMajor) Imag() mat.Mutable { return MutableImag(A) }
