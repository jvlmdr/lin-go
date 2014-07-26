package blas

import (
	"fmt"

	"github.com/jvlmdr/lin-go/zmat"
	"github.com/jvlmdr/lin-go/zvec"
)

// Computes (alpha A x), with A optionally transposed.
//
// Inputs are unchanged.
func MatTimesVecCmplx(alpha complex128, A zmat.Stride, t Transpose, x zvec.Slice) zvec.Slice {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	y := zvec.MakeSlice(size.Rows)
	MatTimesVecPlusVecCmplxNoCopy(alpha, A, t, x, 0, y)
	return y
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls zgemv.
//
// Inputs are unchanged.
func MatTimesVecPlusVecCmplx(alpha complex128, A zmat.Stride, t Transpose, x zvec.Slice, beta complex128, y zvec.Const) zvec.Slice {
	z := zvec.MakeSliceCopy(y)
	MatTimesVecPlusVecCmplxNoCopy(alpha, A, t, x, beta, z)
	return z
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls zgemv.
//
// The result is returned in y.
// A and x are unchanged.
func MatTimesVecPlusVecCmplxNoCopy(alpha complex128, A zmat.Stride, t Transpose, x zvec.Slice, beta complex128, y zvec.Slice) {
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

	zgemv(t, zmat.Rows(A), zmat.Cols(A), alpha, A.Elems, A.Stride,
		[]complex128(x), 1, beta, []complex128(y), 1)
}
