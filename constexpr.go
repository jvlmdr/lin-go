package mat

// This file contains Const matrix expressions which do not have a Mutable partner.

import "github.com/jackvalmadre/go-vec"

// Returns an nxn identity matrix.
func Identity(n int) Const {
	return identityExpr{n}
}

type identityExpr struct {
	N int
}

func (expr identityExpr) Size() Size {
	return Size{expr.N, expr.N}
}

func (expr identityExpr) At(i, j int) float64 {
	if i == j {
		return 1
	}
	return 0
}

// Returns an nxn read-only diagonal matrix.
func Diag(v vec.Const) Const {
	return diagExpr{v}
}

type diagExpr struct {
	Vector vec.Const
}

func (expr diagExpr) Size() Size {
	n := expr.Vector.Size()
	return Size{n, n}
}

func (expr diagExpr) At(i, j int) float64 {
	if i == j {
		return expr.Vector.At(i)
	}
	return 0
}

// Returns an mxn zero matrix.
func Zeros(m, n int) Const {
	return zeroExpr{m, n}
}

type zeroExpr struct{ M, N int }

func (expr zeroExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr zeroExpr) At(i, j int) float64 {
	return 0
}

// Returns an mxn one matrix.
func Ones(m, n int) Const {
	return onesExpr{m, n}
}

type onesExpr struct{ M, N int }

func (expr onesExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr onesExpr) At(i, j int) float64 {
	return 1
}

// Returns an mxn constant matrix.
func Constant(m, n int, alpha float64) Const {
	return constantExpr{m, n, alpha}
}

type constantExpr struct {
	M, N  int
	Alpha float64
}

func (expr constantExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr constantExpr) At(i, j int) float64 {
	return expr.Alpha
}
