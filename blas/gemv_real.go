package blas

import (
	"fmt"

	"github.com/jvlmdr/lin-go/mat"
	"github.com/jvlmdr/lin-go/vec"
)

// Computes (alpha A x), with A optionally transposed.
//
// Inputs are unchanged.
func MatTimesVec(alpha float64, A mat.Stride, t Transpose, x vec.Slice) vec.Slice {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	y := vec.MakeSlice(size.Rows)
	MatTimesVecPlusVecNoCopy(alpha, A, t, x, 0, y)
	return y
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls dgemv.
//
// Inputs are unchanged.
func MatTimesVecPlusVec(alpha float64, A mat.Stride, t Transpose, x vec.Slice, beta float64, y vec.Const) vec.Slice {
	z := vec.MakeSliceCopy(y)
	MatTimesVecPlusVecNoCopy(alpha, A, t, x, beta, z)
	return z
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls dgemv.
//
// The result is returned in y.
// A and x are unchanged.
func MatTimesVecPlusVecNoCopy(alpha float64, A mat.Stride, t Transpose, x vec.Slice, beta float64, y vec.Slice) {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	if size.Cols != x.Len() {
		panic(fmt.Sprintf("A and x have incompatible dimension (%v and %v)", size, x.Len()))
	}
	if size.Rows != y.Len() {
		panic(fmt.Sprintf("A and y have incompatible dimension (%v and %v)", size, y.Len()))
	}

	dgemv(t, mat.Rows(A), mat.Cols(A), alpha, A.Elems, A.Stride,
		[]float64(x), 1, beta, []float64(y), 1)
}
