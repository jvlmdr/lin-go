package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// This file contains operations which address matrices as vectors.

type ConstVecer interface {
	ConstVec() zvec.Const
}

type MutableVecer interface {
	MutableVec() zvec.Mutable
}

// Vectorizes a constant matrix in column-major order.
// Calls A.ConstVec() if A is a ConstVecer.
func Vec(A Const) zvec.Const {
	if A, ok := A.(ConstVecer); ok {
		return A.ConstVec()
	}
	return vecExpr{A}
}

type vecExpr struct{ Matrix Const }

func (x vecExpr) Len() int { return x.Matrix.Size().Area() }

func (x vecExpr) At(i int) complex128 {
	M := x.Matrix
	rows := Rows(M)
	p := i % rows
	q := i / rows
	return M.At(p, q)
}

// Vectorizes a mutable matrix in column-major order.
// Calls A.MutableVec() if A is a MutableVecer.
func MutableVec(A Mutable) zvec.Mutable {
	if A, ok := A.(MutableVecer); ok {
		return A.MutableVec()
	}
	return mutableVecExpr{A}
}

type mutableVecExpr struct{ Matrix Mutable }

func (x mutableVecExpr) Len() int { return x.Matrix.Size().Area() }

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

type ConstColer interface {
	ConstCol(j int) zvec.Const
}

type MutableColer interface {
	MutableCol(j int) zvec.Mutable
}

// Accesses a column in a constant matrix as a vector.
// Calls A.ConstCol() if A is a ConstColer.
func Col(A Const, j int) zvec.Const {
	if A, ok := A.(ConstColer); ok {
		return A.ConstCol(j)
	}
	return columnExpr{A, j}
}

type columnExpr struct {
	Matrix Const
	J      int
}

func (col columnExpr) Len() int            { return Rows(col.Matrix) }
func (col columnExpr) At(i int) complex128 { return col.Matrix.At(i, col.J) }

// Accesses a column in a mutable matrix as a vector.
// Calls A.MutableCol() if A is a MutableColer.
func MutableCol(A Mutable, j int) zvec.Mutable {
	if A, ok := A.(MutableColer); ok {
		return A.MutableCol(j)
	}
	return mutableColExpr{A, j}
}

type mutableColExpr struct {
	Matrix Mutable
	J      int
}

func (col mutableColExpr) Len() int                { return Rows(col.Matrix) }
func (col mutableColExpr) At(i int) complex128     { return col.Matrix.At(i, col.J) }
func (col mutableColExpr) Set(i int, v complex128) { col.Matrix.Set(i, col.J, v) }

type ConstRower interface {
	ConstRow(i int) zvec.Const
}

type MutableRower interface {
	MutableRow(i int) zvec.Mutable
}

// Accesses a row in a constant matrix as a vector.
// Calls A.ConstRow() if A is a ConstRower.
func Row(A Const, i int) zvec.Const {
	if A, ok := A.(ConstRower); ok {
		return A.ConstRow(i)
	}
	return rowExpr{A, i}
}

type rowExpr struct {
	Matrix Const
	I      int
}

func (row rowExpr) Len() int            { return Rows(row.Matrix) }
func (row rowExpr) At(j int) complex128 { return row.Matrix.At(row.I, j) }

// Accesses a row in a constant matrix as a vector.
// Calls A.MutableRow() if A is a MutableRower.
func MutableRow(A Mutable, i int) zvec.Mutable {
	if A, ok := A.(MutableRower); ok {
		return A.MutableRow(i)
	}
	return mutableRowExpr{A, i}
}

type mutableRowExpr struct {
	Matrix Mutable
	I      int
}

func (row mutableRowExpr) Len() int                { return Rows(row.Matrix) }
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

func (expr mutableDiagVecExpr) Len() int {
	rows, cols := RowsCols(expr.Matrix)
	return min(rows, cols)
}

func (expr mutableDiagVecExpr) At(i int) complex128 {
	return expr.Matrix.At(i, i)
}

func (expr mutableDiagVecExpr) Set(i int, x complex128) {
	expr.Matrix.Set(i, i, x)
}
