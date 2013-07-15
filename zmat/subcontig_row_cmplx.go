package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContiguousRowMajorSubmat) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContiguousRowMajorSubmat) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContiguousRowMajorSubmat) Imag() mat.Mutable { return MutableImag(A) }
