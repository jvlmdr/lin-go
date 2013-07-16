package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Computes (alpha A x), with A optionally transposed.
//
// Inputs are unchanged.
func MatrixTimesVector(alpha float64, A mat.SemiContiguousColMajor, t Transpose, x vec.Slice) vec.Slice {
	size := A.Size()
	if t != NoTrans {
		size = size.T()
	}

	y := vec.Make(size.Rows)
	MatrixTimesVectorPlusVectorInPlace(alpha, A, t, x, 0, y)
	return y
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls DGEMV.
//
// Inputs are unchanged.
func MatrixTimesVectorPlusVector(alpha float64, A mat.SemiContiguousColMajor, t Transpose, x vec.Slice, beta float64, y vec.Const) vec.Slice {
	z := vec.MakeCopy(y)
	MatrixTimesVectorPlusVectorInPlace(alpha, A, t, x, beta, z)
	return z
}

// Computes (alpha A x + beta y), with A optionally transposed.
//
// Calls DGEMV.
//
// The result is returned in y.
// A and x are unchanged.
func MatrixTimesVectorPlusVectorInPlace(alpha float64, A mat.SemiContiguousColMajor, t Transpose, x vec.Slice, beta float64, y vec.Slice) {
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

	DGEMV(t, mat.Rows(A), mat.Cols(A), alpha, A.ColMajorArray(), A.Stride(),
		[]float64(x), 1, beta, []float64(y), 1)
}
