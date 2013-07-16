package lapack

import (
	"fmt"
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
func SolveComplexSquareInPlace(A zmat.ColMajor, b zvec.Slice) ComplexLU {
	if zmat.Rows(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := zmat.Contiguous{b.Size(), 1, []complex128(b)}
	lu := SolveComplexSquareMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
func SolveComplexSquareMatrix(A zmat.Const, B zmat.Const) (zmat.Contiguous, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	X := zmat.MakeContiguousCopy(B)
	lu := SolveComplexSquareMatrixInPlace(Q, X)
	return X, lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
//
// Result is returned in B.
func SolveComplexSquareMatrixInPlace(A zmat.ColMajor, B zmat.ColMajor) ComplexLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrix dimensions are incompatible")
	}

	n := zmat.Rows(A)
	ipiv := make(IntList, n)

	info := ZGESV(zmat.Rows(A), zmat.Cols(B), A.ColMajorArray(), A.Stride(), ipiv,
		B.ColMajorArray(), B.Stride())
	if info != 0 {
		panic(fmt.Sprintf("info was non-zero (%d)", info))
	}

	return ComplexLU{A, ipiv}
}
