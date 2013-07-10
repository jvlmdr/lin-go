package mat

// This file contains Const matrix expressions which do not have a Mutable partner.

import "github.com/jackvalmadre/go-vec"

// Returns an nxn identity matrix.
func Identity(n int) IdentityExpr {
	return IdentityExpr{n}
}

type IdentityExpr struct {
	N int
}

func (expr IdentityExpr) Size() Size {
	return expr.N, expr.N
}

func (expr IdentityExpr) At(i, j int) float64 {
	if i == j {
		return 1
	}
	return 0
}

// Returns an nxn read-only diagonal matrix.
func Diag(v vec.Const) DiagExpr {
	return DiagExpr{v}
}

type DiagExpr struct {
	Vector vec.Const
}

func (expr DiagExpr) Size() Size {
	n := expr.Vector.Size()
	return Size{n, n}
}

func (expr DiagExpr) At(i, j int) float64 {
	if i == j {
		return expr.Vector.At(i)
	}
	return 0
}

// Returns an mxn zero matrix.
func Zeros(m, n int) ZeroExpr {
	return ZeroExpr{m, n}
}

type ZeroExpr struct{ M, N int }

func (expr ZeroExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr ZeroExpr) At(i, j int) float64 {
	return 0
}

// Returns an mxn one matrix.
func Ones(m, n int) OnesExpr {
	return OnesExpr{m, n}
}

type OnesExpr struct{ M, N int }

func (expr OnesExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr OnesExpr) At(i, j int) float64 {
	return 1
}

// Returns an mxn constant matrix.
func Constant(m, n int, alpha float64) ConstantExpr {
	return ConstantExpr{m, n, alpha}
}

type ConstantExpr struct {
	M, N  int
	Alpha float64
}

func (expr ConstantExpr) Size() Size {
	return Size{expr.M, expr.N}
}

func (expr ConstantExpr) At(i, j int) float64 {
	return expr.Alpha
}
