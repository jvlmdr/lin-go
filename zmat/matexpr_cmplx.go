package zmat

// This file contains operations which modify the shape of a matrix.

import "math/cmplx"

// Returns a thin wrapper for the conjugate (or Hermitian) transpose of a constant matrix.
func H(A Const) Const { return hExpr{A} }

type hExpr struct{ Matrix Const }

func (expr hExpr) Size() Size { return expr.Matrix.Size().T() }

func (expr hExpr) At(i, j int) complex128 {
	return cmplx.Conj(expr.Matrix.At(j, i))
}

// Returns a thin wrapper for the conjugate (or Hermitian) transpose of a mutable matrix.
func MutableH(A Mutable) Mutable { return mutableHExpr{A} }

type mutableHExpr struct{ Matrix Mutable }

func (expr mutableHExpr) Size() Size { return expr.Matrix.Size().T() }

func (expr mutableHExpr) At(i, j int) complex128 {
	return cmplx.Conj(expr.Matrix.At(j, i))
}

func (expr mutableHExpr) Set(i, j int, v complex128) {
	expr.Matrix.Set(j, i, cmplx.Conj(v))
}
