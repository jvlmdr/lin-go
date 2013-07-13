package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"unsafe"
)

// #cgo linux LDFLAGS: -llapack -lblas
// #cgo darwin LDFLAGS: -framework Accelerate
// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves a square system of equations A x = b.
//
// Copies A and b, then calls SolveInPlace.
func Solve(A mat.Const, b vec.Const) vec.Slice {
	// A x = b becomes Q x = u.
	Q := mat.MakeContiguousCopy(A)
	u := vec.MakeSliceCopy(b)
	SolveInPlace(Q, u)
	return u
}

// Solves a square system of equations A x = b.
// A is over-written with the LU factorization, result is returned in b.
//
// Wraps b as a matrix, then calls SolveMatInPlace.
func SolveInPlace(A mat.SubContiguous, b vec.Slice) {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	SolveMatInPlace(A, B)
}

// Calls SolveMatInPlace on copies of A and B.
func SolveMat(A mat.Const, B mat.Const) mat.Contiguous {
	// A X = B becomes Q X = U.
	Q := mat.MakeContiguousCopy(A)
	U := mat.MakeContiguousCopy(B)
	SolveMatInPlace(Q, U)
	return U
}

//	/* Subroutine */ int dgetrf_(integer *m, integer *n, doublereal *a, integer *
//		lda, integer *ipiv, integer *info);

// A must be square.
// B must be sub-contiguous, column-major.
//
// Result returned in B.
// A will be over-written.
func SolveMatInPlace(A mat.SubContiguous, B mat.SubContiguous) {
	// Compute LU factorization.
	lu := LUInPlace(A)
	// Solve system.
	lu.SolveInPlace(B)
}

type LU struct {
	Matrix mat.SubContiguous
	Ipiv   []C.integer
}

// Calls DGETRF.
// A must be square.
func LUInPlace(A mat.SubContiguous) LU {
	// Factorize A.
	rows, cols := mat.RowsCols(A)

	m := C.integer(rows)
	n := C.integer(cols)
	a := (*C.doublereal)(unsafe.Pointer(&A.Array()[0]))
	lda := C.integer(A.Stride())
	ipiv := make([]C.integer, min(rows, cols))
	var info C.integer

	C.dgetrf_(&m, &n, a, &lda, &ipiv[0], &info)

	return LU{A, ipiv}
}

//	/* Subroutine */ int dgetrs_(char *trans, integer *n, integer *nrhs, 
//		doublereal *a, integer *lda, integer *ipiv, doublereal *b, integer *
//		ldb, integer *info);

// Calls DGETRS.
// B must be sub-contiguous, column-major.
func (lu LU) SolveInPlace(B mat.SubContiguous) {
	if B.RowMajor() {
		panic("B is not column major")
	}

	trans := C.char('N')
	if lu.Matrix.RowMajor() {
		trans = C.char('T')
	}
	n := C.integer(mat.Rows(lu.Matrix))
	nrhs := C.integer(mat.Cols(B))
	a := (*C.doublereal)(unsafe.Pointer(&lu.Matrix.Array()[0]))
	lda := C.integer(lu.Matrix.Stride())
	b := (*C.doublereal)(unsafe.Pointer(&B.Array()[0]))
	ldb := C.integer(B.Stride())
	var info C.integer

	C.dgetrs_(&trans, &n, &nrhs, a, &lda, &lu.Ipiv[0], b, &ldb, &info)
}
