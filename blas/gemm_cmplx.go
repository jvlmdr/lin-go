package blas

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/zmat"
)

// Computes (alpha A B), with A and B optionally transposed.
//
// Inputs are unchanged.
func ComplexMatrixTimesMatrix(alpha complex128, A zmat.ColMajor, tA Transpose, B zmat.ColMajor, tB Transpose) zmat.Contiguous {
	// Get sizes of A and B after transposing.
	sizeA := A.Size()
	if tA != NoTrans {
		sizeA = sizeA.T()
	}
	sizeB := B.Size()
	if tB != NoTrans {
		sizeB = sizeB.T()
	}

	C := zmat.Make(sizeA.Rows, sizeB.Cols)
	ComplexMatrixTimesMatrixPlusMatrixInPlace(alpha, A, tA, B, tB, 0, C)
	return C
}

// Computes (alpha A B + C), with A and B optionally transposed.
//
// Calls ZGEMM.
//
// Inputs are unchanged.
func ComplexMatrixTimesMatrixPlusMatrix(alpha complex128, A zmat.ColMajor, tA Transpose, B zmat.ColMajor, tB Transpose, C zmat.Const) zmat.Contiguous {
	D := zmat.MakeCopy(C)
	ComplexMatrixTimesMatrixPlusMatrixInPlace(alpha, A, tA, B, tB, 1, D)
	return D
}

// Computes (alpha A B + beta C), with A and B optionally transposed.
//
// Calls ZGEMM.
//
// The result is returned in C.
// A and B are unchanged.
func ComplexMatrixTimesMatrixPlusMatrixInPlace(alpha complex128, A zmat.ColMajor, tA Transpose, B zmat.ColMajor, tB Transpose, beta complex128, C zmat.ColMajor) {
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

	ZGEMM(tA, tB, sizeA.Rows, sizeB.Cols, sizeA.Cols,
		alpha, A.ColMajorArray(), A.Stride(), B.ColMajorArray(), B.Stride(),
		beta, C.ColMajorArray(), C.Stride())
}
