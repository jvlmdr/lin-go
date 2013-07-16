package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A X = B where A is full rank.
//
// Calls DGELS.
//
// B will contain the solution.
// A will be over-written with either the LQ or QR factorization.
func SolveFullRankMatrixInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, B mat.SemiContiguousColMajor) {
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < max(mat.Rows(A), mat.Cols(A)) {
		m, n := mat.RowsCols(A)
		// Transpose dimensions if necessary.
		if trans != NoTrans {
			m, n = n, m
		}
		if mat.Rows(B) < m {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}

	trans_ := C.char(trans)
	m := C.integer(mat.Rows(A))
	n := C.integer(mat.Cols(A))
	nrhs := C.integer(mat.Cols(B))
	p_a := (*C.doublereal)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_b := (*C.doublereal)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	var info C.integer

	// Determine optimal workspace size.
	work := make([]float64, 1)
	p_work := (*C.doublereal)(unsafe.Pointer(&work[0]))
	lwork := C.integer(-1)

	C.dgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &lwork, &info)

	// Allocate optimal workspace size.
	lwork = C.integer(int(forceToReal(work[0])))
	p_work = nil
	if int(lwork) > 0 {
		work = make([]float64, int(lwork))
		p_work = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}

	// Solve system.
	C.dgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &lwork, &info)
}
