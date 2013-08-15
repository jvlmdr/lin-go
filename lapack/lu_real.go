package lapack

import "github.com/jackvalmadre/lin-go/mat"

// Describes an LU factorization.
type LUFact struct {
	A    mat.Stride
	Ipiv IntList
}

// Like LUNoCopy() except A is left intact.
func LU(A mat.Const) (LUFact, error) {
	return LUNoCopy(mat.MakeStrideCopy(A))
}

// Computes an LU factorization.
// Calls dgetrf.
// This is essentially the factorization used by SolveSquare (through dgesv).
//
// Note that A does not need to be square to compute the decomposition.
// However, it does need to be square to call Solve().
//
// The input A will be over-written.
func LUNoCopy(A mat.Stride) (LUFact, error) {
	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, mat.Rows(A))

	info := dgetrf(A.Rows, A.Cols, A.Elems, A.Stride, ipiv)
	if info != 0 {
		return LUFact{}, ErrNonZeroInfo
	}
	lu := LUFact{A, ipiv}
	return lu, nil
}
