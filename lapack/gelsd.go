package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

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

////////////////////////////////////////////////////////////////////////////////
// COMPLEX

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Returns solution, rank, singular values.
//
// rcond is used to determine the effective rank of A.
// Singular values <= rcond * sigma_1 are considered zero.
// If rcond < 0, machine precision is used.
func SolveComplex(A zmat.Const, b zvec.Const, rcond float64) (zvec.Slice, int, []float64) {
	if zmat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := zmat.RowsCols(A)
	Q := zmat.MakeContiguousCopy(A)
	// Allocate enough space for constraints or solution.
	ux := zvec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	x := ux.Subvec(0, n)
	zvec.Copy(u, b)

	rank, sigma := SolveComplexInPlace(Q, ux, rcond)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveComplexInPlace(A zmat.SemiContiguousColMajor, b zvec.Slice, rcond float64) (int, []float64) {
	B := zmat.ContiguousColMajor{b.Size(), 1, []complex128(b)}
	return SolveComplexMatrixInPlace(A, B, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls ZGELSD.
//
// Returns solution, rank and singular values.
func SolveComplexMatrix(A zmat.Const, B zmat.Const, rcond float64) (zmat.SemiContiguousColMajor, int, []float64) {
	if zmat.Rows(A) != zmat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)
	// Translate into Q X = U.
	Q := zmat.MakeContiguousCopy(A)
	UX := zmat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(zmat.MakeRect(0, 0, m, nrhs))
	X := UX.Submat(zmat.MakeRect(0, 0, n, nrhs))
	zmat.Copy(U, B)

	rank, sigma := SolveComplexMatrixInPlace(Q, UX, rcond)
	return X, rank, sigma
}
