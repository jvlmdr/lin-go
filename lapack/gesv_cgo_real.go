package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A X = B where A is square.
//
// Calls DGESV.
//
// Result is returned in B.
func SquareSolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor) RealLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrix dimensions are incompatible")
	}

	n := C.integer(mat.Rows(A))
	nrhs := C.integer(mat.Cols(B))
	p_A := (*C.doublereal)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_B := (*C.doublereal)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	var info C.integer
	ipiv := make([]C.integer, int(n))

	C.dgesv_(&n, &nrhs, p_A, &lda, &ipiv[0], p_B, &ldb, &info)

	return RealLU{A, ipiv}
}
