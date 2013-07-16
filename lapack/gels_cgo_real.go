package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"unsafe"
)

// #include "../f2c.h"
// #include "../clapack.h"
import "C"

// Solves A X = B where A is full rank.
//
// Calls DGELS.
//
// B must be large enough to hold both the constraints and the solution (not simultaneously).
// Returns a matrix which references the elements of B.
// A will be over-written with either the LQ or QR factorization.
func SolveFullRankMatrixInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, B mat.SemiContiguousColMajor) mat.SemiContiguousColMajor {
	// Transpose dimensions if necessary.
	numeqs, numvars := mat.RowsCols(A)
	if trans != NoTrans {
		numeqs, numvars = numvars, numeqs
	}
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < max(mat.Rows(A), mat.Cols(A)) {
		if mat.Rows(B) < numeqs {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}
	X := mat.SemiContiguousSubmat(B, mat.MakeRect(0, 0, numvars, mat.Cols(B)))

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

	return X
}
