package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A x = b where A is square.
//
// Calls DGESV.
func SquareSolve(A mat.Const, b vec.Const) (vec.Slice, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	x := vec.MakeSliceCopy(b)
	lu := SquareSolveInPlace(Q, x)
	return x, lu
}

// Solves A x = b where A is square.
//
// Calls DGESV.
//
// Result is returned in b.
func SquareSolveInPlace(A mat.SemiContiguousColMajor, b vec.Slice) RealLU {
	if mat.Cols(A) != b.Size() {
		panic("Matrix and vector dimensions are incompatible")
	}
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	lu := SquareSolveMatrixInPlace(A, B)
	return lu
}

// Solves A X = B where A is square.
//
// Calls DGESV.
func SquareSolveMatrix(A mat.Const, B mat.Const) (mat.ContiguousColMajor, RealLU) {
	Q := mat.MakeContiguousCopy(A)
	X := mat.MakeContiguousCopy(B)
	lu := SquareSolveMatrixInPlace(Q, X)
	return X, lu
}

// Solves A X = B where A is square.
//
// Calls DGESV.
//
// Result is returned in B.
func SquareSolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor) RealLU {
	if !A.Size().Square() {
		panic("System of equations is not square")
	}
	if mat.Cols(A) != mat.Rows(B) {
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
