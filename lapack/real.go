package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"unsafe"
)

// #include "f2c.h"
// #include "clapack.h"
import "C"

// Describes an LU factorization.
type LU struct {
	A    mat.SemiContiguousColMajor
	Ipiv []C.integer
}

////////////////////////////////////////////////////////////////////////////////
// DGESV

// Solves A x = b where A is square.
//
// Calls DGESV.
func SquareSolve(A mat.Const, b vec.Const) (vec.Slice, LU) {
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
func SquareSolveInPlace(A mat.SemiContiguousColMajor, b vec.Slice) LU {
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
func SquareSolveMatrix(A mat.Const, B mat.Const) (mat.ContiguousColMajor, LU) {
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
func SquareSolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor) LU {
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

	return LU{A, ipiv}
}

////////////////////////////////////////////////////////////////////////////////
// DGELS

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func FullRankSolve(A mat.Const, b vec.Const) vec.Slice {
	if mat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := mat.RowsCols(A)
	Q := mat.MakeContiguousCopy(A)
	// Allocate enough space for input and solution.
	ux := vec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	vec.Copy(u, b)

	FullRankSolveInPlace(Q, NoTrans, ux)
	return x
}

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func FullRankSolveInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, b vec.Slice) {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	FullRankSolveMatrixInPlace(A, trans, B)
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
func FullRankSolveMatrix(A mat.Const, B mat.Const) mat.SemiContiguousColMajor {
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	// Translate into Q X = U.
	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)
	Q := mat.MakeContiguousCopy(A)
	// Allocate enough space for constraints and solution.
	UX := mat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(mat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(mat.MakeRect(0, 0, n, nrhs))
	mat.Copy(U, B)

	FullRankSolveMatrixInPlace(Q, NoTrans, UX)
	return X
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
//
// B will contain the solution.
// A will be over-written with either the LQ or QR factorization.
func FullRankSolveMatrixInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, B mat.SemiContiguousColMajor) {
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
	C_lwork := C.integer(-1)
	C.dgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)

	// Allocate optimal workspace size.
	lwork := int(work[0])
	work = make([]float64, lwork)
	p_work = (*C.doublereal)(unsafe.Pointer(&work[0]))
	C_lwork = C.integer(lwork)

	// Solve system.
	C.dgels_(&trans_, &m, &n, &nrhs, p_a, &lda, p_b, &ldb, p_work, &C_lwork, &info)
}

////////////////////////////////////////////////////////////////////////////////
// DGELSD

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Returns solution, rank, singular values.
//
// rcond is used to determine the effective rank of A.
// Singular values <= rcond * sigma_1 are considered zero.
// If rcond < 0, machine precision is used.
func Solve(A mat.Const, b vec.Const, rcond float64) (vec.Slice, int, []float64) {
	if mat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := mat.RowsCols(A)
	Q := mat.MakeContiguousCopy(A)
	// Allocate enough space for constraints or solution.
	ux := vec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	vec.Copy(u, b)

	rank, sigma := SolveInPlace(Q, ux, rcond)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveInPlace(A mat.SemiContiguousColMajor, b vec.Slice, rcond float64) (int, []float64) {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	return SolveMatrixInPlace(A, B, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Returns solution, rank and singular values.
func SolveMatrix(A mat.Const, B mat.Const, rcond float64) (mat.SemiContiguousColMajor, int, []float64) {
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)
	// Translate into Q X = U.
	Q := mat.MakeContiguousCopy(A)
	UX := mat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(mat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(mat.MakeRect(0, 0, n, nrhs))
	mat.Copy(U, B)

	rank, sigma := SolveMatrixInPlace(Q, UX, rcond)
	return X, rank, sigma
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Result is returned in B, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveMatrixInPlace(A mat.SemiContiguousColMajor, B mat.SemiContiguousColMajor, rcond float64) (int, []float64) {
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < max(mat.Rows(A), mat.Cols(A)) {
		if mat.Rows(B) < mat.Rows(A) {
			panic("Not enough rows to contain constraints")
		} else {
			panic("Not enough rows to contain solution")
		}
	}

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
	liwork := C.integer(-1)

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, &iwork[0], &info)

	// Allocate optimal workspace size.
	// Floating-point workspace.
	lwork = C.integer(C.doublereal(work[0]))
	work = make([]float64, int(lwork))
	p_work = (*C.doublereal)(unsafe.Pointer(&work[0]))
	// Integer workspace.
	liwork = iwork[0]
	iwork = make([]C.integer, int(liwork))

	C.dgelsd_(&m, &n, &nrhs, p_A, &lda, p_B, &ldb, p_s, &rcond_, &rank,
		p_work, &lwork, &iwork[0], &info)

	return int(rank), sigma
}
