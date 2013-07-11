package zmat

import "github.com/jackvalmadre/go-vec/zvec"

// This file contains operations which treat matrices as vectors.

// Adds two matrices of the same dimension.
// Returns a thin wrapper which evaluates the operation on demand.
func Plus(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Unvec(zvec.Plus(Vec(A), Vec(B)), rows, cols)
}

// Subtracts one matrix  of the same dimension.
// Returns a thin wrapper which evaluates the operation on demand.
func Minus(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Unvec(zvec.Minus(Vec(A), Vec(B)), rows, cols)
}

// Scales all elements of a matrix.
// Returns a thin wrapper which evaluates the operation on demand.
func Scale(k complex128, A Const) Const {
	rows, cols := RowsCols(A)
	return Unvec(zvec.Scale(k, Vec(A)), rows, cols)
}

// Element-wise multiplication of two matrices.
// Returns a thin wrapper which evaluates the operation on demand.
func VectorMultiply(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Unvec(zvec.Multiply(Vec(A), Vec(B)), rows, cols)
}
