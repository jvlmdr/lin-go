package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContiguousSubmat) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContiguousSubmat) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContiguousSubmat) Imag() mat.Mutable { return MutableImag(A) }
