package lapack

import "github.com/jackvalmadre/lin-go/zmat"

// Describes an LU factorization.
type LUFactCmplx struct {
	A    zmat.Stride
	Ipiv IntList
}

// Like LUCmplxNoCopy() except A is left intact.
func LUCmplx(A zmat.Const) (LUFactCmplx, error) {
	return LUCmplxNoCopy(zmat.MakeStrideCopy(A))
}

// Computes an LU factorization.
// Calls zgetrf.
// This is essentially the factorization used by SolveSquare (through zgesv).
//
// Note that A does not need to be square to compute the decomposition.
// However, it does need to be square to call Solve().
//
// The input A will be over-written.
func LUCmplxNoCopy(A zmat.Stride) (LUFactCmplx, error) {
	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, zmat.Rows(A))

	info := zgetrf(A.Rows, A.Cols, A.Elems, A.Stride, ipiv)
	if info != 0 {
		return LUFactCmplx{}, ErrNonZeroInfo
	}
	lu := LUFactCmplx{A, ipiv}
	return lu, nil
}
