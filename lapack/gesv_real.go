package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is square.
//
// Calls DGESV.
func SquareSolve(A mat.Const, b vec.Const) (vec.Slice, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	x := vec.MakeSliceCopy(b)
	lu := SquareSolveInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls DGESV.
//
// Result is returned in b.
func SquareSolveInPlace(A mat.SemiContiguousColMajor, b vec.Slice) RealLU {
	if mat.Rows(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	lu := SquareSolveMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls DGESV.
func SquareSolveMatrix(A mat.Const, B mat.Const) (mat.ContiguousColMajor, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	X := mat.MakeContiguousCopy(B)
	lu := SquareSolveMatrixInPlace(Q, X)
	return X, lu
}
