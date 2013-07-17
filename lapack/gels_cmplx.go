package lapack

// ZGELS functions which are trivially mapped to ZGELS functions.

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func SolveComplexFullRank(A zmat.Const, b zvec.Const) zvec.Slice {
	if zmat.Rows(A) != b.Len() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for input and solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	zvec.Copy(u, b)

	return SolveComplexFullRankInPlace(Q, NoTrans, ux)
}

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func SolveComplexFullRankInPlace(A zmat.ColMajor, trans Transpose, b zvec.Slice) zvec.Slice {
	B := zmat.Contiguous{b.Len(), 1, []complex128(b)}
	X := SolveComplexFullRankMatrixInPlace(A, trans, B)
	return zmat.ContiguousCol(X, 0)
}

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
func SolveComplexFullRankMatrix(A zmat.Const, B zmat.Const) zmat.ColMajor {
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
	zmat.Copy(U, B)
	return SolveComplexFullRankMatrixInPlace(Q, NoTrans, UX)
}

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
//
// B must be large enough to hold both the constraints and the solution (not simultaneously).
// Returns a matrix which references the elements of B.
// A will be over-written with either the LQ or QR factorization.
func SolveComplexFullRankMatrixInPlace(A zmat.ColMajor, trans Transpose, B zmat.ColMajor) zmat.ColMajor {
	size := A.Size()
	// Transpose dimensions if necessary.
	if trans != NoTrans {
		size = size.T()
	}
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if zmat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	ZGELSAuto(trans, zmat.Rows(A), zmat.Cols(A), zmat.Cols(B),
		A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride())

	return zmat.ColMajorSubmat(B, zmat.MakeRect(0, 0, size.Cols, zmat.Cols(B)))
}

// Automatically allocates workspace.
func ZGELSAuto(trans Transpose, m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) (info int) {
	var (
		lwork = -1
		work  = make([]complex128, 1)
	)
	info = ZGELS(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
	if info != 0 {
		return
	}

	lwork = int(forceToReal(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}
	return ZGELS(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
}
