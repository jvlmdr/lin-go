package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Solves A x = b where A is square.
//
// Calls ZGESV.
func SquareSolveComplex(A zmat.Const, b zvec.Const) (zvec.Slice, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	x := zvec.MakeSliceCopy(b)
	lu := SquareSolveComplexInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls ZGESV.
//
// Result is returned in b.
func SquareSolveComplexInPlace(A zmat.SemiContiguousColMajor, b zvec.Slice) ComplexLU {
	if zmat.Rows(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	lu := SquareSolveComplexMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
func SquareSolveComplexMatrix(A zmat.Const, B zmat.Const) (zmat.ContiguousColMajor, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	X := zmat.MakeContiguousCopy(B)
	lu := SquareSolveComplexMatrixInPlace(Q, X)
	return X, lu
}
