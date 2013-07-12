package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A SubContiguousColMajor) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A SubContiguousColMajor) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A SubContiguousColMajor) Imag() mat.Mutable { return MutableImag(A) }
