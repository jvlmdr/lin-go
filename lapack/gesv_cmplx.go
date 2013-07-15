package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A x = b where A is square.
//
// Calls ZGESV.
func SquareSolveComplex(A zmat.Const, b zvec.Const) (zvec.Slice, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	x := zvec.MakeSliceCopy(b)
	lu := SquareSolveComplexInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls ZGESV.
//
// Result is returned in b.
func SquareSolveComplexInPlace(A zmat.SemiContiguousColMajor, b zvec.Slice) ComplexLU {
	if zmat.Cols(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	lu := SquareSolveComplexMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
func SquareSolveComplexMatrix(A zmat.Const, B zmat.Const) (zmat.ContiguousColMajor, ComplexLU) {
	Q := zmat.MakeContiguousCopy(A)
	X := zmat.MakeContiguousCopy(B)
	lu := SquareSolveComplexMatrixInPlace(Q, X)
	return X, lu
}

// Solves A X = B where A is square.
//
// Calls ZGESV.
//
// Result is returned in B.
func SquareSolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, B zmat.SemiContiguousColMajor) ComplexLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if zmat.Cols(A) != zmat.Rows(B) {
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
