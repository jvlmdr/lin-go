package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
//
// B will contain the solution.
// A will be over-written with either the LQ or QR factorization.
func SolveComplexFullRankMatrixSystemInPlace(A zmat.SemiContiguousColMajor, trans TransposeMode, B zmat.SemiContiguousColMajor) {
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
	lwork := C.integer(-1)

	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &lwork, &info)

	// Allocate optimal workspace size.
	lwork = C.integer(int(forceToReal(work[0])))
	p_work = nil
	if int(lwork) > 0 {
		work = make([]complex128, int(lwork))
		p_work = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	}

	// Solve system.
	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &lwork, &info)
}
