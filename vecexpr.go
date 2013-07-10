package mat

import "github.com/jackvalmadre/go-vec"

// Adds two matrices. Lazily evaluated.
//
// Provided as a convenience to allow
//	C := mat.DenseCopy(mat.Plus(A, B))
// instead of
//	C := mat.MakeDense(Rows(A), Cols(A))
//	vec.Copy(C.Vec(), vec.Plus(mat.Vec(A), mat.Vec(B)))
func Plus(A, B Const) Const {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return Reshape(vec.Plus(Vec(A), Vec(B)), rows, cols)
}

// Lazily subtracts one matrix from another using default vectorization.
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
