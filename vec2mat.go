package mat

import "github.com/jackvalmadre/go-vec"

// Address a const vector as a column matrix.
type Mat struct{ vec.Const }

func (A Mat) Size() Size {
	return Size{A.Const.Size(), 1}
}

func (A Mat) At(i, j int) float64 {
	return A.Const.At(i)
}

// Address a mutable vector as a column matrix.
type MutableMat struct{ vec.Mutable }

func (A MutableMat) Size() Size {
	return Size{A.Mutable.Size(), 1}
}

func (A MutableMat) At(i, j int) float64 {
	return A.Mutable.At(i)
}

func (A MutableMat) Set(i, j int, v float64) {
	A.Mutable.Set(i, v)
}

// Address a constant vector as a matrix.
func Reshape(x vec.Const, rows, cols int) ReshapeExpr {
	return ReshapeExpr{x, rows, cols}
}

// Gives a constant vector a matrix shape.
type ReshapeExpr struct {
	Vector vec.Const
	Rows   int
	Cols   int
}

func (expr ReshapeExpr) Size() Size {
	return Size{expr.Rows, expr.Cols}
}

func (expr ReshapeExpr) At(i, j int) float64 {
	return expr.Vector.At(i*expr.Cols + j)
}

// Address a mutable vector as a matrix.
func MutableReshape(x vec.Mutable, rows, cols int) MutableReshapeExpr {
	return MutableReshapeExpr{x, rows, cols}
}

// Gives a mutable vector a matrix shape.
type MutableReshapeExpr struct {
	Vector vec.Mutable
	Rows   int
	Cols   int
}

func (expr MutableReshapeExpr) Size() Size {
	return Size{expr.Rows, expr.Cols}
}

func (expr MutableReshapeExpr) At(i, j int) float64 {
	return expr.Vector.At(i*expr.Cols + j)
}

func (expr MutableReshapeExpr) Set(i, j int, v float64) {
	expr.Vector.Set(i*expr.Cols+j, v)
}
