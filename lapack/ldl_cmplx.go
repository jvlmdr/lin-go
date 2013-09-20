package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"fmt"
)

// Describes an LDL factorization (Cholesky with a diagonal matrix).
type LDLFactCmplx struct {
	UpLo UpLo
	A    zmat.Stride
	Ipiv IntList
}

// Like LDLCmplxNoCopy() except A is left intact.
func LDLCmplx(A zmat.Const) (LDLFactCmplx, error) {
	return LDLCmplxNoCopy(zmat.MakeStrideCopy(A))
}

// Computes an LDL factorization.
// Calls zsytrf.
// This is essentially the factorization used by SolveSymmXxx (through zsysv).
//
// The input A will be over-written.
func LDLCmplxNoCopy(A zmat.Stride) (LDLFactCmplx, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}

	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, A.Rows)

	info := zsytrfAuto(uplo, A.Rows, A.Elems, A.Stride, ipiv)
	if info != 0 {
		return LDLFactCmplx{}, ErrNonZeroInfo(info)
	}
	ldl := LDLFactCmplx{uplo, A, ipiv}
	return ldl, nil
}

// Automatically allocates workspace.
func zsytrfAuto(uplo UpLo, n int, a []complex128, lda int, ipiv IntList) int {
	var (
		lwork = -1
		work  = make([]complex128, 1)
	)
	if info := zsytrf(uplo, n, a, lda, ipiv, work, lwork); info != 0 {
		return info
	}

	lwork = int(re(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}
	return zsytrf(uplo, n, a, lda, ipiv, work, lwork)
}

// Vector version of SolveMat().
// Like SolveNoCopy() except B is left intact.
func (ldl LDLFactCmplx) Solve(b zvec.Const) (zvec.Slice, error) {
	return ldl.SolveNoCopy(zvec.MakeSliceCopy(b))
}

// Like SolveMat() except B is left intact.
func (ldl LDLFactCmplx) SolveMat(B zmat.Const) (zmat.Stride, error) {
	return ldl.SolveMatNoCopy(zmat.MakeStrideCopy(B))
}

// Vector version of SolveMatNoCopy().
func (ldl LDLFactCmplx) SolveNoCopy(b zvec.Slice) (zvec.Slice, error) {
	X, err := ldl.SolveMatNoCopy(zmat.StrideMat(b))
	if err != nil {
		return zvec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves a square, symmetric system given its LDL factorization.
// Calls zsytrs.
func (ldl LDLFactCmplx) SolveMatNoCopy(B zmat.Stride) (zmat.Stride, error) {
	// Check that B has the same number of rows as A.
	if ldl.A.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", ldl.A.Size(), B.Size()))
	}

	info := zsytrs(ldl.UpLo, ldl.A.Rows, B.Cols, ldl.A.Elems, ldl.A.Stride, ldl.Ipiv, B.Elems, B.Stride)
	if info != 0 {
		return zmat.Stride{}, ErrNonZeroInfo(info)
	}
	return B, nil
}
