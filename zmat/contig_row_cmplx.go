package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContiguousRowMajor) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContiguousRowMajor) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContiguousRowMajor) Imag() mat.Mutable { return MutableImag(A) }
