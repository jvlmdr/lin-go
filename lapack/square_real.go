package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
)

// Vector version of SolveSquareMat().
// Like SolveSquareNoCopy() except A and b are left intact.
func SolveSquare(A mat.Const, b vec.Const) (vec.Slice, error) {
	return SolveSquareNoCopy(mat.MakeStrideCopy(A), vec.MakeSliceCopy(b))
}

// Like SolveSquareMatNoCopy() except A and B are left intact.
func SolveSquareMat(A mat.Const, B mat.Const) (mat.Stride, error) {
	return SolveSquareMatNoCopy(mat.MakeStrideCopy(A), mat.MakeStrideCopy(B))
}

// Vector version of SolveSquareMatNoCopy().
func SolveSquareNoCopy(A mat.Stride, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	X, err := SolveSquareMat(A, B)
	if err != nil {
		return nil, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is square and full-rank.
//
// Calls dgesv.
//
// Overwrites A and b.
// Result references elements of b.
func SolveSquareMatNoCopy(A mat.Stride, B mat.Stride) (mat.Stride, error) {
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}
	if mat.Rows(A) != mat.Rows(B) {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", A.Size(), B.Size()))
	}

	ipiv := make(IntList, A.Rows)
	info := dgesv(A.Rows, B.Cols, A.Elems, A.Stride, ipiv, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo
	}
	return B, nil
}
