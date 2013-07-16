package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is square.
//
// Calls DGESV.
func SolveSquare(A mat.Const, b vec.Const) (vec.Slice, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	x := vec.MakeSliceCopy(b)
	lu := SolveSquareInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls DGESV.
//
// Result is returned in b.
func SolveSquareInPlace(A mat.SemiContiguousColMajor, b vec.Slice) RealLU {
	if mat.Rows(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	lu := SolveSquareMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls DGESV.
func SolveSquareMatrix(A mat.Const, B mat.Const) (mat.ContiguousColMajor, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	X := mat.MakeContiguousCopy(B)
	lu := SolveSquareMatrixInPlace(Q, X)
	return X, lu
}
