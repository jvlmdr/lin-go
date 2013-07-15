package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplex(A zmat.Const, b zvec.Const) zvec.Slice {
	if zmat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for input and solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	zvec.Copy(u, b)

	FullRankSolveComplexInPlace(Q, NoTrans, ux)
	return x
}

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplexInPlace(A zmat.SemiContiguousColMajor, trans TransposeMode, b zvec.Slice) {
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	FullRankSolveComplexMatrixInPlace(A, trans, B)
}

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplexMatrix(A zmat.Const, B zmat.Const) zmat.SemiContiguousColMajor {
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	// Translate into Q X = U.
	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for constraints and solution.
	UX := zmat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(zmat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(zmat.MakeRect(0, 0, n, nrhs))
	zmat.Copy(U, B)

	FullRankSolveComplexMatrixInPlace(Q, NoTrans, UX)
	return X
}
