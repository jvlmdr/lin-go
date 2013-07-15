package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Returns solution, rank, singular values.
//
// rcond is used to determine the effective rank of A.
// Singular values <= rcond * sigma_1 are considered zero.
// If rcond < 0, machine precision is used.
func Solve(A mat.Const, b vec.Const, rcond float64) (vec.Slice, int, []float64) {
	if mat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := mat.RowsCols(A)
	Q := mat.MakeContiguousCopy(A)
	// Allocate enough space for constraints or solution.
	ux := vec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	vec.Copy(u, b)

	rank, sigma := SolveInPlace(Q, ux, rcond)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveInPlace(A mat.SemiContiguousColMajor, b vec.Slice, rcond float64) (int, []float64) {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	return SolveMatrixInPlace(A, B, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Returns solution, rank and singular values.
func SolveMatrix(A mat.Const, B mat.Const, rcond float64) (mat.SemiContiguousColMajor, int, []float64) {
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)
	// Translate into Q X = U.
	Q := mat.MakeContiguousCopy(A)
	UX := mat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(mat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(mat.MakeRect(0, 0, n, nrhs))
	mat.Copy(U, B)

	rank, sigma := SolveMatrixInPlace(Q, UX, rcond)
	return X, rank, sigma
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Result is returned in B, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor, rcond float64) (int, []float64) {
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < max(mat.Rows(A), mat.Cols(A)) {
		if mat.Rows(B) < mat.Rows(A) {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}

	m := C.integer(mat.Rows(A))
	n := C.integer(mat.Cols(A))
	nrhs := C.integer(mat.Cols(B))
	p_A := (*C.doublereal)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_B := (*C.doublereal)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	rcond_ := C.doublereal(rcond)
	var rank C.integer
	var info C.integer

	// Singular values.
	sigma := make([]float64, min(mat.Rows(A), mat.Cols(A)))
	p_s := (*C.doublereal)(unsafe.Pointer(&sigma[0]))

	// Determine optimal workspace sizes.
	// Floating-point workspace.
	work := make([]float64, 1)
	p_work := (*C.doublereal)(unsafe.Pointer(&work[0]))
	lwork := C.integer(-1)
	// Integer workspace.
	iwork := make([]C.integer, 1)
	liwork := C.integer(-1)

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, &iwork[0], &info)

	// Allocate optimal workspace size.
	// Floating-point workspace.
	lwork = C.integer(C.doublereal(work[0]))
	work = make([]float64, int(lwork))
	p_work = (*C.doublereal)(unsafe.Pointer(&work[0]))
	// Integer workspace.
	liwork = iwork[0]
	iwork = make([]C.integer, int(liwork))

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, &iwork[0], &info)

	return int(rank), sigma
}

////////////////////////////////////////////////////////////////////////////////
// COMPLEX

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
