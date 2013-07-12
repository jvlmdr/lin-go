package zmat

import "github.com/jackvalmadre/go-vec/zvec"

// This file contains operations which address vectors as matrices.

// Addresses a const vector as a column matrix.
func Mat(x zvec.Const) Const { return matExpr{x} }

type matExpr struct{ Vector zvec.Const }

func (A matExpr) Size() Size             { return Size{A.Vector.Size(), 1} }
func (A matExpr) At(i, j int) complex128 { return A.Vector.At(i) }

// Addresses a mutable vector as a column matrix.
func MutableMat(x zvec.Mutable) Mutable { return mutableMatExpr{x} }

type mutableMatExpr struct{ Vector zvec.Mutable }

func (A mutableMatExpr) Size() Size                 { return Size{A.Vector.Size(), 1} }
func (A mutableMatExpr) At(i, j int) complex128     { return A.Vector.At(i) }
func (A mutableMatExpr) Set(i, j int, v complex128) { A.Vector.Set(i, v) }

// Address a constant vector as a matrix.
func Unvec(x zvec.Const, rows, cols int) Const { return unvecExpr{x, rows, cols} }

// Gives a constant vector a matrix shape.
type unvecExpr struct {
	Vector zvec.Const
	Rows   int
	Cols   int
}

func (expr unvecExpr) Size() Size { return Size{expr.Rows, expr.Cols} }

func (expr unvecExpr) At(i, j int) complex128 {
	return expr.Vector.At(i*expr.Cols + j)
}

// Address a mutable vector as a matrix.
func MutableUnvec(x zvec.Mutable, rows, cols int) Mutable {
	return mutableUnvecExpr{x, rows, cols}
}

// Gives a mutable vector a matrix shape.
type mutableUnvecExpr struct {
	Vector zvec.Mutable
	Rows   int
	Cols   int
}

func (expr mutableUnvecExpr) Size() Size { return Size{expr.Rows, expr.Cols} }

func (expr mutableUnvecExpr) At(i, j int) complex128 {
	return expr.Vector.At(i*expr.Cols + j)
}

func (expr mutableUnvecExpr) Set(i, j int, v complex128) {
	expr.Vector.Set(i*expr.Cols+j, v)
}
