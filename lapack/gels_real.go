package lapack

// DGELS functions which are trivially mapped to ZGELS functions.

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func SolveFullRank(A mat.Const, b vec.Const) vec.Slice {
	if mat.Rows(A) != b.Size() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := mat.RowsCols(A)
	Q := mat.MakeContiguousCopy(A)
	// Allocate enough space for input and solution.
	ux := vec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	vec.Copy(u, b)

	return SolveFullRankInPlace(Q, NoTrans, ux)
}

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func SolveFullRankInPlace(A mat.ColMajor, trans Transpose, b vec.Slice) vec.Slice {
	B := mat.Contiguous{b.Size(), 1, []float64(b)}
	X := SolveFullRankMatrixInPlace(A, trans, B)
	return mat.ContiguousCol(X, 0)
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
func SolveFullRankMatrix(A mat.Const, B mat.Const) mat.ColMajor {
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
	mat.Copy(U, B)
	return SolveFullRankMatrixInPlace(Q, NoTrans, UX)
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
//
// B must be large enough to hold both the constraints and the solution (not simultaneously).
// Returns a matrix which references the elements of B.
// A will be over-written with either the LQ or QR factorization.
func SolveFullRankMatrixInPlace(A mat.ColMajor, trans Transpose, B mat.ColMajor) mat.ColMajor {
	size := A.Size()
	// Transpose dimensions if necessary.
	if trans != NoTrans {
		size = size.T()
	}
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if mat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	DGELSAuto(trans, mat.Rows(A), mat.Cols(A), mat.Cols(B),
		A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride())

	return mat.ColMajorSubmat(B, mat.MakeRect(0, 0, size.Cols, mat.Cols(B)))
}

// Automatically allocates workspace.
func DGELSAuto(trans Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int) (info int) {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info = DGELS(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
	if info != 0 {
		return
	}

	lwork = int(forceToReal(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return DGELS(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
}
