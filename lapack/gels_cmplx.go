package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplex(A zmat.Const, b zvec.Const) zvec.Slice {
	if zmat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for input and solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	zvec.Copy(u, b)

	FullRankSolveComplexInPlace(Q, NoTrans, ux)
	return x
}

// Solves A x = b where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplexInPlace(A zmat.SemiContiguousColMajor, trans TransposeMode, b zvec.Slice) {
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	FullRankSolveComplexMatrixInPlace(A, trans, B)
}

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
func FullRankSolveComplexMatrix(A zmat.Const, B zmat.Const) zmat.SemiContiguousColMajor {
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	// Translate into Q X = U.
	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for constraints and solution.
	UX := zmat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(zmat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(zmat.MakeRect(0, 0, n, nrhs))
	zmat.Copy(U, B)

	FullRankSolveComplexMatrixInPlace(Q, NoTrans, UX)
	return X
}

// Solves A X = B where A is full rank.
//
// Calls ZGELS.
//
// B will contain the solution.
// A will be over-written with either the LQ or QR factorization.
func FullRankSolveComplexMatrixInPlace(A zmat.SemiContiguousColMajor, trans TransposeMode, B zmat.SemiContiguousColMajor) {
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
	C_lwork := C.integer(-1)
	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)

	// Allocate optimal workspace size.
	lwork := int(forceToReal(work[0]))
	work = make([]complex128, lwork)
	p_work = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	C_lwork = C.integer(lwork)

	// Solve system.
	C.zgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)
}
