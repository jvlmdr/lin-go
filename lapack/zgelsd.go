package lapack

// ZGELSD functions which use CGo.

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Returns solution, rank, singular values.
//
// rcond is used to determine the effective rank of A.
// Singular values <= rcond * sigma_1 are considered zero.
// If rcond < 0, machine precision is used.
func SolveComplex(A zmat.Const, b zvec.Const, rcond float64) (zvec.Slice, int, []float64) {
	if zmat.Rows(A) != b.Len() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for constraints or solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	zvec.Copy(u, b)

	return SolveComplexInPlace(Q, ux, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexInPlace(A zmat.ColMajor, b zvec.Slice, rcond float64) (zvec.Slice, int, []float64) {
	B := zmat.Contiguous{b.Len(), 1, []complex128(b)}
	X, rank, sigma := SolveComplexMatrixInPlace(A, B, rcond)
	x := zmat.ContiguousCol(X, 0)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Returns solution, rank and singular values.
func SolveComplexMatrix(A zmat.Const, B zmat.Const, rcond float64) (zmat.ColMajor, int, []float64) {
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)
	// Translate into Q X = U.
	Q := zmat.MakeContiguousCopy(A)
	UX := zmat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(zmat.MakeRect(0, 0, m, nrhs))
	zmat.Copy(U, B)

	return SolveComplexMatrixInPlace(Q, UX, rcond)
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Result is returned in B, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexMatrixInPlace(A zmat.ColMajor, B zmat.ColMajor, rcond float64) (zmat.ColMajor, int, []float64) {
	size := A.Size()
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if zmat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	sigma := make([]float64, min(size.Rows, size.Cols))

	rank, info := ZGELSDAuto(zmat.Rows(A), zmat.Cols(A), zmat.Cols(B),
		A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride(), sigma, rcond)
	if info != 0 {
		panic(fmt.Sprintf("info was non-zero (%d)", info))
	}

	X := zmat.ColMajorSubmat(B, zmat.MakeRect(0, 0, size.Cols, zmat.Cols(B)))
	return X, rank, sigma
}

// Automatically allocates workspace.
func ZGELSDAuto(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64) (rank int, info int) {
	var (
		lwork = -1
		work  = make([]complex128, 1)
		rwork = make([]float64, 1)
		iwork = make(IntList, 1)
	)
	rank, info = ZGELSD(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, rwork, iwork)
	if info != 0 {
		return
	}

	lwork = int(real(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}

	lrwork := int(rwork[0])
	rwork = nil
	if lrwork > 0 {
		rwork = make([]float64, lrwork)
	}

	liwork := int(iwork[0])
	iwork = nil
	if liwork > 0 {
		iwork = make(IntList, liwork)
	}

	return ZGELSD(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, rwork, iwork)
}
