package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A X = B where A is square.
//
// Calls ZGESV.
//
// Result is returned in B.
func SquareSolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, B zmat.SemiContiguousColMajor) ComplexLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrix dimensions are incompatible")
	}

	n := C.integer(zmat.Rows(A))
	nrhs := C.integer(zmat.Cols(B))
	p_A := (*C.doublecomplex)(unsafe.Pointer(&A.ColMajorArray()[0]))
	lda := C.integer(A.Stride())
	p_B := (*C.doublecomplex)(unsafe.Pointer(&B.ColMajorArray()[0]))
	ldb := C.integer(B.Stride())
	var info C.integer
	ipiv := make([]C.integer, int(n))

	C.zgesv_(&n, &nrhs, p_A, &lda, &ipiv[0], p_B, &ldb, &info)

	return ComplexLU{A, ipiv}
}
