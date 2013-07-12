package zmat

// This file contains operations which modify the shape of a matrix.

import (
	"github.com/jackvalmadre/lin-go/mat"
	"math/cmplx"
)

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

// Returns a thin wrapper for the real part of a constant matrix.
func Real(A Const) mat.Const { return realExpr{A} }

type realExpr struct{ Matrix Const }

func (expr realExpr) Size() mat.Size { return mat.Size(expr.Matrix.Size()) }

func (expr realExpr) At(i, j int) float64 {
	return real(expr.Matrix.At(i, j))
}

// Returns a thin wrapper for the real part of a mutable matrix.
func MutableReal(A Mutable) mat.Mutable { return mutableRealExpr{A} }

type mutableRealExpr struct{ Matrix Mutable }

func (expr mutableRealExpr) Size() mat.Size { return mat.Size(expr.Matrix.Size()) }

func (expr mutableRealExpr) At(i, j int) float64 {
	return real(expr.Matrix.At(i, j))
}

func (expr mutableRealExpr) Set(i, j int, u float64) {
	v := imag(expr.Matrix.At(i, j))
	expr.Matrix.Set(i, j, complex(u, v))
}

// Returns a thin wrapper for the imaginary part of a constant matrix.
func Imag(A Const) mat.Const { return imagExpr{A} }

type imagExpr struct{ Matrix Const }

func (expr imagExpr) Size() mat.Size { return mat.Size(expr.Matrix.Size()) }

func (expr imagExpr) At(i, j int) float64 {
	return imag(expr.Matrix.At(i, j))
}

// Returns a thin wrapper for the imaginary part of a mutable matrix.
func MutableImag(A Mutable) mat.Mutable { return mutableImagExpr{A} }

type mutableImagExpr struct{ Matrix Mutable }

func (expr mutableImagExpr) Size() mat.Size { return mat.Size(expr.Matrix.Size()) }

func (expr mutableImagExpr) At(i, j int) float64 {
	return imag(expr.Matrix.At(i, j))
}

func (expr mutableImagExpr) Set(i, j int, v float64) {
	u := real(expr.Matrix.At(i, j))
	expr.Matrix.Set(i, j, complex(u, v))
}

// Returns a thin wrapper for A + Bi.
func Complex(A, B mat.Const) Const { return complexExpr{A, B} }

type complexExpr struct{ Real, Imag mat.Const }

func (expr complexExpr) Size() Size { return Size(expr.Real.Size()) }

func (expr complexExpr) At(i, j int) complex128 {
	return complex(expr.Real.At(i, j), expr.Imag.At(i, j))
}

// Returns a thin wrapper for A + Bi.
func MutableComplex(A, B mat.Mutable) Mutable { return mutableComplexExpr{A, B} }

type mutableComplexExpr struct{ Real, Imag mat.Mutable }

func (expr mutableComplexExpr) Size() Size { return Size(expr.Real.Size()) }

func (expr mutableComplexExpr) At(i, j int) complex128 {
	return complex(expr.Real.At(i, j), expr.Imag.At(i, j))
}

func (expr mutableComplexExpr) Set(i, j int, x complex128) {
	expr.Real.Set(i, j, real(x))
	expr.Imag.Set(i, j, imag(x))
}
