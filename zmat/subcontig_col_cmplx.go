package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContiguousColMajorSubmat) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContiguousColMajorSubmat) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContiguousColMajorSubmat) Imag() mat.Mutable { return MutableImag(A) }
