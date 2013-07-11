package zmat

// This file contains operations which modify the shape of a matrix.

// Returns a thin wrapper which transposes a constant matrix.
func T(A Const) Const { return tExpr{A} }

type tExpr struct{ Matrix Const }

func (expr tExpr) Size() Size             { return expr.Matrix.Size().T() }
func (expr tExpr) At(i, j int) complex128 { return expr.Matrix.At(j, i) }

// Returns a thin wrapper which transposes a mutable matrix.
func MutableT(A Mutable) Mutable { return mutableTExpr{A} }

type mutableTExpr struct{ Matrix Mutable }

func (expr mutableTExpr) Size() Size                 { return expr.Matrix.Size().T() }
func (expr mutableTExpr) At(i, j int) complex128     { return expr.Matrix.At(j, i) }
func (expr mutableTExpr) Set(i, j int, v complex128) { expr.Matrix.Set(j, i, v) }

// Returns a thin wrapper which selects a constant submatrix.
func Submatrix(A Const, r Rect) Const { return submatrixExpr{A, r} }

type submatrixExpr struct {
	Matrix Const
	Rect   Rect
}

func (expr submatrixExpr) Size() Size { return expr.Rect.Size() }

func (expr submatrixExpr) At(i, j int) complex128 {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	return expr.Matrix.At(i-i0, j-j0)
}

// Returns a thin wrapper which selects a mutable submatrix.
func MutableSubmatrix(A Mutable, r Rect) Mutable {
	return mutableSubmatrixExpr{A, r}
}

type mutableSubmatrixExpr struct {
	Matrix Mutable
	Rect   Rect
}

func (expr mutableSubmatrixExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr mutableSubmatrixExpr) At(i, j int) complex128 {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	return expr.Matrix.At(i-i0, j-j0)
}

func (expr mutableSubmatrixExpr) Set(i, j int, x complex128) {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	expr.Matrix.Set(i-i0, j-j0, x)
}

// Address a constant matrix by a different shape.
// Equivalent to vectorizing then unvectorizing.
func Reshape(A Const, rows, cols int) Const { return Unvec(Vec(A), rows, cols) }

// Address a mutable matrix by a different shape.
// Equivalent to vectorizing then unvectorizing.
func MutableReshape(A Mutable, rows, cols int) Mutable {
	return MutableUnvec(MutableVec(A), rows, cols)
}
