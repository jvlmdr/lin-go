package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
)

// Computes (alpha A B), with A and B optionally transposed.
//
// Inputs are unchanged.
func MatTimesMat(alpha float64, A mat.Stride, tA Transpose, B mat.Stride, tB Transpose) mat.Stride {
	// Get sizes of A and B after transposing.
	sizeA := A.Size()
	if tA != NoTrans {
		sizeA = sizeA.T()
	}
	sizeB := B.Size()
	if tB != NoTrans {
		sizeB = sizeB.T()
	}

	C := mat.MakeStride(sizeA.Rows, sizeB.Cols)
	MatTimesMatPlusMatNoCopy(alpha, A, tA, B, tB, 0, C)
	return C
}

// Computes (alpha A B + C), with A and B optionally transposed.
//
// Calls dgemm.
//
// Inputs are unchanged.
func MatTimesMatPlusMat(alpha float64, A mat.Stride, tA Transpose, B mat.Stride, tB Transpose, C mat.Const) mat.Stride {
	D := mat.MakeStrideCopy(C)
	MatTimesMatPlusMatNoCopy(alpha, A, tA, B, tB, 1, D)
	return D
}

// Computes (alpha A B + beta C), with A and B optionally transposed.
//
// Calls dgemm.
//
// The result is returned in C.
// A and B are unchanged.
func MatTimesMatPlusMatNoCopy(alpha float64, A mat.Stride, tA Transpose, B mat.Stride, tB Transpose, beta float64, C mat.Stride) {
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
	if sizeA.Rows != mat.Rows(C) {
		panic(fmt.Sprintf("A and C have incompatible dimensions (%v and %v)", sizeA, C.Size()))
	}
	if sizeB.Cols != mat.Cols(C) {
		panic(fmt.Sprintf("B and C have incompatible dimensions (%v and %v)", sizeB, C.Size()))
	}

	dgemm(tA, tB, sizeA.Rows, sizeB.Cols, sizeA.Cols, alpha, A.Elems, A.Stride,
		B.Elems, B.Stride, beta, C.Elems, C.Stride)
}
