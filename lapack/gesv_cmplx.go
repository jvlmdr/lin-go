package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Solves A x = b where A is square.
//
// Calls ZGESV.
func SolveComplexSquare(A zmat.Const, b zvec.Const) (zvec.Slice, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	x := zvec.MakeSliceCopy(b)
	lu := SolveComplexSquareInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls ZGESV.
//
// Result is returned in b.
func SolveComplexSquareInPlace(A zmat.SemiContiguousColMajor, b zvec.Slice) ComplexLU {
	if zmat.Rows(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	lu := SolveComplexSquareMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
func SolveComplexSquareMatrix(A zmat.Const, B zmat.Const) (zmat.ContiguousColMajor, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	X := zmat.MakeContiguousCopy(B)
	lu := SolveComplexSquareMatrixInPlace(Q, X)
	return X, lu
}
