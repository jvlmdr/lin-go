package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
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
	vec.Copy(u, b)

	return SolveInPlace(Q, ux, rcond)
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// The result is written to b, which must be big enough to hold constraints and solution (not simultaneously).
// Returns rank and singular values.
func SolveInPlace(A mat.ColMajor, b vec.Slice, rcond float64) (vec.Slice, int, []float64) {
	B := mat.Contiguous{b.Size(), 1, []float64(b)}
	X, rank, sigma := SolveMatrixInPlace(A, B, rcond)
	x := mat.ContiguousCol(X, 0)
	return x, rank, sigma
}

// Solves A x = b where A is not necessarily full rank.
//
// Calls DGELSD.
//
// Returns solution, rank and singular values.
func SolveMatrix(A mat.Const, B mat.Const, rcond float64) (mat.ColMajor, int, []float64) {
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)
	// Translate into Q X = U.
	Q := mat.MakeContiguousCopy(A)
	UX := mat.MakeContiguous(max(m, n), nrhs)
	U := UX.Submat(mat.MakeRect(0, 0, m, nrhs))
	mat.Copy(U, B)

	return SolveMatrixInPlace(Q, UX, rcond)
}

// Solves A X = B where A is not necessarily full rank.
//
// Calls DGELSD.
//
// B must be big enough to hold both constraints and solution (not simultaneously).
// Returns solution, rank and singular values.
// Solution references same data as B.
func SolveMatrixInPlace(A mat.ColMajor, B mat.ColMajor, rcond float64) (mat.ColMajor, int, []float64) {
	size := A.Size()
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if mat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	sigma := make([]float64, min(size.Rows, size.Cols))

	rank, info := DGELSDAuto(mat.Rows(A), mat.Cols(A), mat.Cols(B),
		A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride(), sigma, rcond)
	if info != 0 {
		panic(fmt.Sprintf("info was non-zero (%d)", info))
	}

	X := mat.SemiContiguousSubmat(B, mat.MakeRect(0, 0, size.Cols, mat.Cols(B)))
	return X, rank, sigma
}

// Automatically allocates workspace.
func DGELSDAuto(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64) (rank int, info int) {
	var (
		lwork = -1
		work  = make([]float64, 1)
		iwork = make(IntList, 1)
	)
	rank, info = DGELSD(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, iwork)
	if info != 0 {
		return
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}

	liwork := int(iwork[0])
	iwork = nil
	if liwork > 0 {
		iwork = make(IntList, liwork)
	}

	return DGELSD(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, iwork)
}
