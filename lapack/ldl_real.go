package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
)

// Describes an LDL factorization (Cholesky with a diagonal matrix).
type LDLFact struct {
	UpLo UpLo
	A    mat.Stride
	Ipiv IntList
}

// Like LDLNoCopy() except A is left intact.
func LDL(A mat.Const) (LDLFact, error) {
	return LDLNoCopy(mat.MakeStrideCopy(A))
}

// Computes an LDL factorization.
// Calls dsytrf.
// This is essentially the factorization used by SolveSymmXxx (through dsysv).
//
// The input A will be over-written.
func LDLNoCopy(A mat.Stride) (LDLFact, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}

	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, A.Rows)

	info := dsytrfAuto(uplo, A.Rows, A.Elems, A.Stride, ipiv)
	if info != 0 {
		return LDLFact{}, ErrNonZeroInfo
	}
	ldl := LDLFact{uplo, A, ipiv}
	return ldl, nil
}

// Automatically allocates workspace.
func dsytrfAuto(uplo UpLo, n int, a []float64, lda int, ipiv IntList) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	if info := dsytrf(uplo, n, a, lda, ipiv, work, lwork); info != 0 {
		return info
	}

	lwork = int(re(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dsytrf(uplo, n, a, lda, ipiv, work, lwork)
}

// Vector version of SolveMat().
// Like SolveNoCopy() except B is left intact.
func (ldl LDLFact) Solve(b vec.Const) (vec.Slice, error) {
	return ldl.SolveNoCopy(vec.MakeSliceCopy(b))
}

// Like SolveMat() except B is left intact.
func (ldl LDLFact) SolveMat(B mat.Const) (mat.Stride, error) {
	return ldl.SolveMatNoCopy(mat.MakeStrideCopy(B))
}

// Vector version of SolveMatNoCopy().
func (ldl LDLFact) SolveNoCopy(b vec.Slice) (vec.Slice, error) {
	X, err := ldl.SolveMatNoCopy(mat.StrideMat(b))
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves a square, symmetric system given its LDL factorization.
// Calls dsytrs.
func (ldl LDLFact) SolveMatNoCopy(B mat.Stride) (mat.Stride, error) {
	// Check that B has the same number of rows as A.
	if ldl.A.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", ldl.A.Size(), B.Size()))
	}

	info := dsytrs(ldl.UpLo, ldl.A.Rows, B.Cols, ldl.A.Elems, ldl.A.Stride, ldl.Ipiv, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo
	}
	return B, nil
}
