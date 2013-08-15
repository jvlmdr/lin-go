package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/zmat"
)

// Computes (alpha A B), with A and B optionally transposed.
//
// Inputs are unchanged.
func MatTimesMatCmplx(alpha complex128, A zmat.Stride, tA Transpose, B zmat.Stride, tB Transpose) zmat.Stride {
	// Get sizes of A and B after transposing.
	sizeA := A.Size()
	if tA != NoTrans {
		sizeA = sizeA.T()
	}
	sizeB := B.Size()
	if tB != NoTrans {
		sizeB = sizeB.T()
	}

	C := zmat.MakeStride(sizeA.Rows, sizeB.Cols)
	MatTimesMatPlusMatCmplxNoCopy(alpha, A, tA, B, tB, 0, C)
	return C
}

// Computes (alpha A B + C), with A and B optionally transposed.
//
// Calls zgemm.
//
// Inputs are unchanged.
func MatTimesMatPlusMatCmplx(alpha complex128, A zmat.Stride, tA Transpose, B zmat.Stride, tB Transpose, C zmat.Const) zmat.Stride {
	D := zmat.MakeStrideCopy(C)
	MatTimesMatPlusMatCmplxNoCopy(alpha, A, tA, B, tB, 1, D)
	return D
}

// Computes (alpha A B + beta C), with A and B optionally transposed.
//
// Calls zgemm.
//
// The result is returned in C.
// A and B are unchanged.
func MatTimesMatPlusMatCmplxNoCopy(alpha complex128, A zmat.Stride, tA Transpose, B zmat.Stride, tB Transpose, beta complex128, C zmat.Stride) {
	// Get sizes of A and B after transposing.
	sizeA := A.Size()
	if tA != NoTrans {
		sizeA = sizeA.T()
	}
	sizeB := B.Size()
	if tB != NoTrans {
		sizeB = sizeB.T()
	}

	if sizeA.Cols != sizeB.Rows {
		panic(fmt.Sprintf("A and B have incompatible dimensions (%v and %v)", sizeA, sizeB))
	}
	if sizeA.Rows != zmat.Rows(C) {
		panic(fmt.Sprintf("A and C have incompatible dimensions (%v and %v)", sizeA, C.Size()))
	}
	if sizeB.Cols != zmat.Cols(C) {
		panic(fmt.Sprintf("B and C have incompatible dimensions (%v and %v)", sizeB, C.Size()))
	}

	zgemm(tA, tB, sizeA.Rows, sizeB.Cols, sizeA.Cols, alpha, A.Elems, A.Stride,
		B.Elems, B.Stride, beta, C.Elems, C.Stride)
}
