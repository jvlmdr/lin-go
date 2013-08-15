package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
)

// Vector version of SolvePosDefMat().
// Like SolvePosDefNoCopy() except A and b are left intact.
func SolvePosDef(A mat.Const, b vec.Const) (vec.Slice, error) {
	C := mat.MakeStrideCopy(A)
	d := vec.MakeSliceCopy(b)
	x, err := SolvePosDefNoCopy(C, d)
	if err != nil {
		return nil, err
	}
	return x, nil
}

// Like SolvePosDefMatNoCopy() except A and B are left intact.
func SolvePosDefMat(A mat.Const, B mat.Const) (mat.Stride, error) {
	C := mat.MakeStrideCopy(A)
	D := mat.MakeStrideCopy(B)
	X, err := SolvePosDefMatNoCopy(C, D)
	if err != nil {
		return mat.Stride{}, err
	}
	return X, nil
}

// Vector version of SolvePosDefMatNoCopy().
func SolvePosDefNoCopy(A mat.Stride, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	if _, err := SolvePosDefMatNoCopy(A, B); err != nil {
		return vec.Slice{}, err
	}
	return b, nil
}

// Solves A X = B where A is symmetric and positive-definite.
//
// Calls dposv.
//
// Overwrites A and B.
// Returns X which references the elements of B.
func SolvePosDefMatNoCopy(A mat.Stride, B mat.Stride) (mat.Stride, error) {
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}
	if mat.Rows(A) != mat.Rows(B) {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", A.Size(), B.Size()))
	}

	const uplo = LowerTriangle
	info := dposv(uplo, A.Rows, B.Cols, A.Elems, A.Stride, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo
	}
	return B, nil
}
