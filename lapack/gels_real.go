package lapack

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
func SolveFullRankInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, b vec.Slice) vec.Slice {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	X := SolveFullRankMatrixInPlace(A, trans, B)
	return mat.ContiguousCol(X, 0)
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
func SolveFullRankMatrix(A mat.Const, B mat.Const) mat.SemiContiguousColMajor {
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
