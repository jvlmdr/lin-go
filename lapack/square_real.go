package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is square.
//
// Calls dgesv.
func SolveSquare(A mat.Const, b vec.Const) (vec.Slice, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	x := vec.MakeSliceCopy(b)
	lu := SolveSquareInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls dgesv.
//
// Result is returned in b.
func SolveSquareInPlace(A mat.ColMajor, b vec.Slice) RealLU {
	if mat.Rows(A) != b.Len() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := mat.Contiguous{b.Len(), 1, []float64(b)}
	lu := SolveNSquareInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls dgesv.
func SolveNSquare(A mat.Const, B mat.Const) (mat.Contiguous, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	X := mat.MakeContiguousCopy(B)
	lu := SolveNSquareInPlace(Q, X)
	return X, lu
}

// Solves A X = B where A is square.
//
// Calls dgesv.
//
// Result is returned in B.
func SolveNSquareInPlace(A mat.ColMajor, B mat.ColMajor) RealLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrix dimensions are incompatible")
	}

	n := mat.Rows(A)
	ipiv := make(IntList, n)

	info := dgesv(mat.Rows(A), mat.Cols(B), A.ColMajorArray(), A.Stride(), ipiv,
		B.ColMajorArray(), B.Stride())
	if info != 0 {
		panic(fmt.Sprintf("info was non-zero (%d)", info))
	}

	return RealLU{A, ipiv}
}
