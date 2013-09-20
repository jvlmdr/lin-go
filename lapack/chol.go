package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Describes a Cholesky factorization.
type CholFact struct {
	A    mat.Stride
	UpLo UpLo
}

// Like CholNoCopy() except A is left intact.
func Chol(A mat.Const) (CholFact, error) {
	return CholNoCopy(mat.MakeStrideCopy(A))
}

// Computes the Cholesky factorization A = L L**T of a symmetric, positive-definite matrix.
// Calls dpotrf.
// This is essentially the factorization used by SolvePosDefXxx (through dposv).
//
// A will be over-written.
func CholNoCopy(A mat.Stride) (CholFact, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic("matrix is not square")
	}

	const uplo = LowerTriangle
	info := dpotrf(uplo, A.Rows, A.Elems, A.Stride)
	if info != 0 {
		return CholFact{}, ErrNonZeroInfo(info)
	}
	ldl := CholFact{A, uplo}
	return ldl, nil
}

// Vector version of SolveMat().
// Like SolveNoCopy() except B is left intact.
func (chol CholFact) Solve(b vec.Const) (vec.Slice, error) {
	return chol.SolveNoCopy(vec.MakeSliceCopy(b))
}

// Like SolveMatNoCopy() except B is left intact.
func (chol CholFact) SolveMat(B mat.Const) (mat.Stride, error) {
	return chol.SolveMatNoCopy(mat.MakeStrideCopy(B))
}

// Vector version of SolveMatNoCopy().
func (chol CholFact) SolveNoCopy(b vec.Slice) (vec.Slice, error) {
	X, err := chol.SolveMatNoCopy(mat.StrideMat(b))
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is symmetric and positive-definite, given its Cholesky decomposition.
func (chol CholFact) SolveMatNoCopy(B mat.Stride) (mat.Stride, error) {
	// Check that B has the same number of rows as A.
	if mat.Rows(chol.A) != mat.Rows(B) {
		panic("numbers of rows do not match")
	}

	info := dpotrs(chol.UpLo, chol.A.Rows, B.Cols, chol.A.Elems, chol.A.Stride, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo(info)
	}
	return B, nil
}
