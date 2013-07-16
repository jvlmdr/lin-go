package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/zmat"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A X = B where A is not necessarily full rank.
//
// Calls DGELSD.
//
// B must be big enough to hold both constraints and solution (not simultaneously).
// Returns solution, rank and singular values.
// Solution references same data as B.
func SolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor, rcond float64) (mat.SemiContiguousColMajor, int, []float64) {
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < max(mat.Rows(A), mat.Cols(A)) {
		if mat.Rows(B) < mat.Rows(A) {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}
	X := mat.SemiContiguousSubmat(B, mat.MakeRect(0, 0, mat.Cols(A), mat.Cols(B)))

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

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, &iwork[0], &info)

	// Allocate optimal workspace size.
	// Floating-point workspace.
	lwork = C.integer(C.doublereal(work[0]))
	p_work = nil
	if int(lwork) > 0 {
		work = make([]float64, int(lwork))
		p_work = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}
	// Integer workspace.
	liwork := int(iwork[0])
	var p_iwork *C.integer
	if liwork > 0 {
		iwork = make([]C.integer, liwork)
		p_iwork = &iwork[0]
	}

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, p_iwork, &info)

	return X, int(rank), sigma
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Result is returned in B, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, B zmat.SemiContiguousColMajor, rcond float64) (zmat.SemiContiguousColMajor, int, []float64) {
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < max(zmat.Rows(A), zmat.Cols(A)) {
		if zmat.Rows(B) < zmat.Rows(A) {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}
	X := zmat.SemiContiguousSubmat(B, zmat.MakeRect(0, 0, zmat.Cols(A), zmat.Cols(B)))

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
	p_work = nil
	if int(lwork) > 0 {
		work = make([]complex128, int(lwork))
		p_work = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	}
	// Real workspace.
	lrwork := int(rwork[0])
	p_rwork = nil
	if lrwork > 0 {
		rwork = make([]float64, lrwork)
		p_rwork = (*C.doublereal)(unsafe.Pointer(&rwork[0]))
	}
	// Integer workspace.
	liwork := int(iwork[0])
	var p_iwork *C.integer
	if liwork > 0 {
		iwork = make([]C.integer, liwork)
		p_iwork = &iwork[0]
	}

	C.zgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, p_rwork, p_iwork, &info)

	return X, int(rank), sigma
}
