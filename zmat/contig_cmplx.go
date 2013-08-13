package zmat

import "github.com/jackvalmadre/lin-go/mat"

// Returns MutableH(A).
func (A Contig) H() Mutable { return MutableH(A) }

// Returns MutableReal(A).
func (A Contig) Real() mat.Mutable { return MutableReal(A) }

// Returns MutableImag(A).
func (A Contig) Imag() mat.Mutable { return MutableImag(A) }
