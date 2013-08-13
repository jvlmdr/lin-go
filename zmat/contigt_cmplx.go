package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A ContigT) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A ContigT) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A ContigT) Imag() mat.Mutable { return MutableImag(A) }
