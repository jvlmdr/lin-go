package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContiguousColMajor) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContiguousColMajor) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContiguousColMajor) Imag() mat.Mutable { return MutableImag(A) }
