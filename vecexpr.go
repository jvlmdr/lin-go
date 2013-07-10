package mat

import "github.com/jackvalmadre/go-vec"

// Adds two matrices of the same dimension.
// Returns a thin wrapper which evaluates the operation on demand.
func Plus(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Reshape(vec.Plus(Vec(A), Vec(B)), rows, cols)
}

// Subtracts one matrix  of the same dimension.
// Returns a thin wrapper which evaluates the operation on demand.
func Minus(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Reshape(vec.Minus(Vec(A), Vec(B)), rows, cols)
}

// Lazily scales a matrix using default vectorization.
func Scale(k float64, A Const) Const {
	rows, cols := RowsCols(A)
	return Reshape(vec.Scale(k, Vec(A)), rows, cols)
}

// Lazy element-wise multiplication of two matrices using default vectorization.
func VectorMultiply(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Reshape(vec.Multiply(Vec(A), Vec(B)), rows, cols)
}
