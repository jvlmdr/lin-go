package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"fmt"
)

// Vector version of SolveMat().
// Like SolveCopy() except b is left intact.
func (lu LUFact) Solve(T bool, b vec.Const) (vec.Slice, error) {
	return lu.SolveNoCopy(T, vec.MakeSliceCopy(b))
}

// Like SolveMatNoCopy() except B is left intact.
func (lu LUFact) SolveMat(T bool, B mat.Const) (mat.Stride, error) {
	return lu.SolveMatNoCopy(T, mat.MakeStrideCopy(B))
}

// Vector version of SolveMatNoCopy().
func (lu LUFact) SolveNoCopy(T bool, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	X, err := lu.SolveMatNoCopy(T, B)
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B (or A**T X = B) where A is square given its LU factorization.
func (lu LUFact) SolveMatNoCopy(T bool, B mat.Stride) (mat.Stride, error) {
	if !lu.A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", lu.A.Size()))
	}

	s := lu.A.Size()
	trans := NoTrans
	if T {
		s = s.T()
		trans = Trans
	}
	if s.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", s, B.Size()))
	}

	info := dgetrs(trans, s.Rows, B.Cols, lu.A.Elems, lu.A.Stride, lu.Ipiv, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo(info)
	}
	return B, nil
}

////////////////////////////////////////////////////////////////////////////////

// Vector version of SolveMat().
// Like SolveCopy() except b is left intact.
func (lu LUFactCmplx) Solve(trans Transpose, b zvec.Const) (zvec.Slice, error) {
	return lu.SolveNoCopy(trans, zvec.MakeSliceCopy(b))
}

// Like SolveMatNoCopy() except B is left intact.
func (lu LUFactCmplx) SolveMat(trans Transpose, B zmat.Const) (zmat.Stride, error) {
	return lu.SolveMatNoCopy(trans, zmat.MakeStrideCopy(B))
}

// Vector version of SolveMatNoCopy().
func (lu LUFactCmplx) SolveNoCopy(trans Transpose, b zvec.Slice) (zvec.Slice, error) {
	B := zmat.StrideMat(b)
	X, err := lu.SolveMatNoCopy(trans, B)
	if err != nil {
		return zvec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is square given its LU factorization.
// Can also solve A**T X = B or A**H X = B.
func (lu LUFactCmplx) SolveMatNoCopy(trans Transpose, B zmat.Stride) (zmat.Stride, error) {
	if !lu.A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", lu.A.Size()))
	}

	s := lu.A.Size()
	if trans != NoTrans {
		s = s.T()
	}
	if s.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", s, B.Size()))
	}

	info := zgetrs(trans, s.Rows, B.Cols, lu.A.Elems, lu.A.Stride, lu.Ipiv, B.Elems, B.Stride)
	if info != 0 {
		return zmat.Stride{}, ErrNonZeroInfo(info)
	}
	return B, nil
}
