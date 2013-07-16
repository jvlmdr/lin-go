package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func SolveFullRankSystem(A mat.Const, b vec.Const) vec.Slice {
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

	SolveFullRankSystemInPlace(Q, NoTrans, ux)
	return x
}

// Solves A x = b where A is full rank.
//
// Calls DGELS.
func SolveFullRankSystemInPlace(A mat.SemiContiguousColMajor, trans TransposeMode, b vec.Slice) {
	B := mat.ContiguousColMajor{b.Size(), 1, []float64(b)}
	SolveFullRankMatrixSystemInPlace(A, trans, B)
}

// Solves A X = B where A is full rank.
//
// Calls DGELS.
func SolveFullRankMatrixSystem(A mat.Const, B mat.Const) mat.SemiContiguousColMajor {
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

	SolveFullRankMatrixSystemInPlace(Q, NoTrans, UX)
	return X
}
