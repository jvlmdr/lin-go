package mat

// This file contains Const matrix expressions which do not have a Mutable partner.

import "github.com/jackvalmadre/lin-go/vec"

// Matrix whose (i, j)-th element is f(i, j).
func MapIndex(m, n int, f func(int, int) float64) Const {
	return mapIndexExpr{m, n, f}
}

type mapIndexExpr struct {
	M int
	N int
	F func(int, int) float64
}

func (expr mapIndexExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr mapIndexExpr) At(i, j int) float64 {
	return expr.F(i, j)
}

// Returns an nxn identity matrix.
func Identity(n int) Const {
	return DiagMat(vec.Ones(n))
}

// Returns an nxn read-only diagonal matrix.
func DiagMat(v vec.Const) Const {
	n := v.Size()
	f := func(i, j int) float64 {
		if i == j {
			return v.At(i)
		}
		return 0
	}
	return MapIndex(n, n, f)
}

// Returns an mxn zero matrix.
func Zeros(m, n int) Const {
	return Constant(m, n, 0)
}

// Returns an mxn one matrix.
func Ones(m, n int) Const {
	return Constant(m, n, 1)
}

// Returns an mxn constant matrix.
func Constant(m, n int, alpha float64) Const {
	return Unvec(vec.Constant(m*n, alpha), m, n)
}

// Returns an mxn constant matrix.
func Randn(m, n int) Const {
	return Unvec(vec.Randn(m*n), m, n)
}
