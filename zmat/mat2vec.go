package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// This file contains operations which address matrices as vectors.

// Vectorizes a constant matrix in column-major order.
func Vec(A Const) zvec.Const { return vecExpr{A} }

type vecExpr struct{ Matrix Const }

func (x vecExpr) Size() int { return x.Matrix.Size().Area() }

func (x vecExpr) At(i int) complex128 {
	M := x.Matrix
	rows := Rows(M)
	p := i % rows
	q := i / rows
	return M.At(p, q)
}

// Vectorizes a mutable matrix in column-major order.
func MutableVec(A Mutable) zvec.Mutable {
	return mutableVecExpr{A}
}

type mutableVecExpr struct{ Matrix Mutable }

func (x mutableVecExpr) Size() int { return x.Matrix.Size().Area() }

func (x mutableVecExpr) At(i int) complex128 {
	M := x.Matrix
	rows := Rows(M)
	p := i % rows
	q := i / rows
	return M.At(p, q)
}

func (x mutableVecExpr) Set(i int, v complex128) {
	M := x.Matrix
	rows := Rows(M)
	p := i % rows
	q := i / rows
	M.Set(p, q, v)
}

// Accesses a column in a constant matrix as a vector.
func Col(A Const, j int) zvec.Const { return columnExpr{A, j} }

type columnExpr struct {
	Matrix Const
	J      int
}

func (col columnExpr) Size() int           { return Rows(col.Matrix) }
func (col columnExpr) At(i int) complex128 { return col.Matrix.At(i, col.J) }

// Accesses a column in a mutable matrix as a vector.
func MutableCol(A Mutable, j int) zvec.Mutable { return mutableColExpr{A, j} }

type mutableColExpr struct {
	Matrix Mutable
	J      int
}

func (col mutableColExpr) Size() int               { return Rows(col.Matrix) }
func (col mutableColExpr) At(i int) complex128     { return col.Matrix.At(i, col.J) }
func (col mutableColExpr) Set(i int, v complex128) { col.Matrix.Set(i, col.J, v) }

// Accesses a row in a constant matrix as a vector.
func Row(A Const, i int) zvec.Const { return rowExpr{A, i} }

type rowExpr struct {
	Matrix Const
	I      int
}

func (row rowExpr) Size() int           { return Rows(row.Matrix) }
func (row rowExpr) At(j int) complex128 { return row.Matrix.At(row.I, j) }

// Accesses a row in a constant matrix as a vector.
func MutableRow(A Mutable, i int) zvec.Mutable { return mutableRowExpr{A, i} }

type mutableRowExpr struct {
	Matrix Mutable
	I      int
}

func (row mutableRowExpr) Size() int               { return Rows(row.Matrix) }
func (row mutableRowExpr) At(j int) complex128     { return row.Matrix.At(row.I, j) }
func (row mutableRowExpr) Set(j int, v complex128) { row.Matrix.Set(row.I, j, v) }

// Returns a constant min(rows, cols)-vector of the leading diagonal.
func DiagVec(A Const) zvec.Const {
	rows, cols := RowsCols(A)
	n := min(rows, cols)
	f := func(i int) complex128 {
		return A.At(i, i)
	}
	return zvec.MapIndex(n, f)
}

// Returns a mutable min(rows, cols)-vector of the leading diagonal.
func MutableDiagVec(A Mutable) zvec.Mutable {
	return mutableDiagVecExpr{A}
}

type mutableDiagVecExpr struct{ Matrix Mutable }

func (expr mutableDiagVecExpr) Size() int {
	rows, cols := RowsCols(expr.Matrix)
	return min(rows, cols)
}

func (expr mutableDiagVecExpr) At(i int) complex128 {
	return expr.Matrix.At(i, i)
}

func (expr mutableDiagVecExpr) Set(i int, x complex128) {
	expr.Matrix.Set(i, i, x)
}
