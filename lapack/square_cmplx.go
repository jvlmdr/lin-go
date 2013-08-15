package lapack

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"fmt"
)

// Vector version of SolveSquareMatCmplx().
// Like SolveSquareCmplxNoCopy() except A and b are left intact.
func SolveSquareCmplx(A zmat.Const, b zvec.Const) (zvec.Slice, error) {
	return SolveSquareCmplxNoCopy(zmat.MakeStrideCopy(A), zvec.MakeSliceCopy(b))
}

// Like SolveSquareMatCmplxNoCopy() except A and B are left intact.
func SolveSquareMatCmplx(A zmat.Const, B zmat.Const) (zmat.Stride, error) {
	return SolveSquareMatCmplxNoCopy(zmat.MakeStrideCopy(A), zmat.MakeStrideCopy(B))
}

// Vector version of SolveSquareMatCmplxNoCopy().
func SolveSquareCmplxNoCopy(A zmat.Stride, b zvec.Slice) (zvec.Slice, error) {
	B := zmat.StrideMat(b)
	X, err := SolveSquareMatCmplx(A, B)
	if err != nil {
		return nil, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is square and full-rank.
//
// Calls zgesv.
//
// Overwrites A and b.
// Result references elements of b.
func SolveSquareMatCmplxNoCopy(A zmat.Stride, B zmat.Stride) (zmat.Stride, error) {
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}
	if zmat.Rows(A) != zmat.Rows(B) {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", A.Size(), B.Size()))
	}

	ipiv := make(IntList, A.Rows)
	info := zgesv(A.Rows, B.Cols, A.Elems, A.Stride, ipiv, B.Elems, B.Stride)
	if info != 0 {
		return zmat.Stride{}, ErrNonZeroInfo
	}
	return B, nil
}
