package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

////////////////////////////////////////////////////////////////////////////////
// ZGESV

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

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
//
// B will contain the solution.
// A will be over-written with either the LQ or QR factorization.
func FullRankSolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, trans TransposeMode, B zmat.SemiContiguousColMajor) {
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < max(zmat.Rows(A), zmat.Cols(A)) {
		m, n := zmat.RowsCols(A)
		// Transpose dimensions if necessary.
		if trans != NoTrans {
			m, n = n, m
		}
		if zmat.Rows(B) < m {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}

	trans_ := C.char(trans)
	m := C.integer(zmat.Rows(A))
	n := C.integer(zmat.Cols(A))
	nrhs := C.integer(zmat.Cols(B))
	p_a := (*C.doublecomplex)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_b := (*C.doublecomplex)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	var info C.integer

	// Determine optimal workspace size.
	work := make([]complex128, 1)
	p_work := (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	C_lwork := C.integer(-1)
	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)

	// Allocate optimal workspace size.
	lwork := int(real(work[0]))
	work = make([]complex128, lwork)
	p_work = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	C_lwork = C.integer(lwork)

	// SolveComplex system.
	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)
}

////////////////////////////////////////////////////////////////////////////////
// ZGELSD

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
	if zmat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for constraints or solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	zvec.Copy(u, b)

	rank, sigma := SolveComplexInPlace(Q, ux, rcond)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexInPlace(A zmat.SemiContiguousColMajor, b zvec.Slice, rcond float64) (int, []float64) {
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	return SolveComplexMatrixInPlace(A, B, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Returns solution, rank and singular values.
func SolveComplexMatrix(A zmat.Const, B zmat.Const, rcond float64) (zmat.SemiContiguousColMajor, int, []float64) {
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)
	// Translate into Q X = U.
	Q := zmat.MakeContiguousCopy(A)
	UX := zmat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(zmat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(zmat.MakeRect(0, 0, n, nrhs))
	zmat.Copy(U, B)

	rank, sigma := SolveComplexMatrixInPlace(Q, UX, rcond)
	return X, rank, sigma
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Result is returned in B, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, B zmat.SemiContiguousColMajor, rcond float64) (int, []float64) {
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < max(zmat.Rows(A), zmat.Cols(A)) {
		if zmat.Rows(B) < zmat.Rows(A) {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}

	m := C.integer(zmat.Rows(A))
	n := C.integer(zmat.Cols(A))
	nrhs := C.integer(zmat.Cols(B))
	p_A := (*C.doublecomplex)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_B := (*C.doublecomplex)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	rcond_ := C.doublereal(rcond)
	var rank C.integer
	var info C.integer

	// Singular values.
	sigma := make([]float64, min(zmat.Rows(A), zmat.Cols(A)))
	p_s := (*C.doublereal)(unsafe.Pointer(&sigma[0]))

	// Determine optimal workspace sizes.
	// Complex workspace.
	lwork := C.integer(-1)
	work := make([]complex128, 1)
	p_work := (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	// Real workspace.
	rwork := make([]float64, 1)
	p_rwork := (*C.doublereal)(unsafe.Pointer(&rwork[0]))
	// Integer workspace.
	iwork := make([]C.integer, 1)

	C.zgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, p_rwork, &iwork[0], &info)

	// Allocate optimal workspace size.
	// Complex workspace.
	lwork = C.integer(int(real(work[0])))
	work = make([]complex128, int(lwork))
	p_work = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	// Real workspace.
	lrwork := int(rwork[0])
	rwork = make([]float64, lrwork)
	p_rwork = (*C.doublereal)(unsafe.Pointer(&rwork[0]))
	// Integer workspace.
	liwork := int(iwork[0])
	iwork = make([]C.integer, liwork)

	C.zgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, p_rwork, &iwork[0], &info)

	return int(rank), sigma
}
