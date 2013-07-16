package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Computes (alpha A x), with A optionally transposed.
//
// Inputs are unchanged.
func ComplexMatrixTimesVector(alpha complex128, A zmat.ColMajor, t Transpose, x zvec.Slice) zvec.Slice {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	y := zvec.Make(size.Rows)
	ComplexMatrixTimesVectorPlusVectorInPlace(alpha, A, t, x, 0, y)
	return y
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls ZGEMV.
//
// Inputs are unchanged.
func ComplexMatrixTimesVectorPlusVector(alpha complex128, A zmat.ColMajor, t Transpose, x zvec.Slice, beta complex128, y zvec.Const) zvec.Slice {
	z := zvec.MakeCopy(y)
	ComplexMatrixTimesVectorPlusVectorInPlace(alpha, A, t, x, beta, z)
	return z
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls ZGEMV.
//
// The result is returned in y.
// A and x are unchanged.
func ComplexMatrixTimesVectorPlusVectorInPlace(alpha complex128, A zmat.ColMajor, t Transpose, x zvec.Slice, beta complex128, y zvec.Slice) {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	if size.Cols != x.Size() {
		panic(fmt.Sprintf("A and x have incompatible dimension (%v and %v)", size, x.Size()))
	}
	if size.Rows != y.Size() {
		panic(fmt.Sprintf("A and y have incompatible dimension (%v and %v)", size, y.Size()))
	}

	ZGEMV(t, zmat.Rows(A), zmat.Cols(A), alpha, A.ColMajorArray(), A.Stride(),
		[]complex128(x), 1, beta, []complex128(y), 1)
}
