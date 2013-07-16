package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
)

// Computes (alpha A B), with A and B optionally transposed.
//
// Inputs are unchanged.
func MatrixTimesMatrix(alpha float64, A mat.ColMajor, tA Transpose, B mat.ColMajor, tB Transpose) mat.Contiguous {
	// Get sizes of A and B after transposing.
	sizeA := A.Size()
	if tA != NoTrans {
		sizeA = sizeA.T()
	}
	sizeB := B.Size()
	if tB != NoTrans {
		sizeB = sizeB.T()
	}

	C := mat.Make(sizeA.Rows, sizeB.Cols)
	MatrixTimesMatrixPlusMatrixInPlace(alpha, A, tA, B, tB, 0, C)
	return C
}

// Computes (alpha A B + C), with A and B optionally transposed.
//
// Calls DGEMM.
//
// Inputs are unchanged.
func MatrixTimesMatrixPlusMatrix(alpha float64, A mat.ColMajor, tA Transpose, B mat.ColMajor, tB Transpose, C mat.Const) mat.Contiguous {
	D := mat.MakeCopy(C)
	MatrixTimesMatrixPlusMatrixInPlace(alpha, A, tA, B, tB, 1, D)
	return D
}

// Computes (alpha A B + beta C), with A and B optionally transposed.
//
// Calls DGEMM.
//
// The result is returned in C.
// A and B are unchanged.
func MatrixTimesMatrixPlusMatrixInPlace(alpha float64, A mat.ColMajor, tA Transpose, B mat.ColMajor, tB Transpose, beta float64, C mat.ColMajor) {
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

	DGEMM(tA, tB, sizeA.Rows, sizeB.Cols, sizeA.Cols,
		alpha, A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride(),
		beta, C.ColMajorArray(), C.Stride())
}
