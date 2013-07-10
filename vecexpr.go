package mat

import "github.com/jackvalmadre/go-vec"

type ContiguousLayout struct {
	Rows     int
	Cols     int
	RowMajor bool
}

func (c ContiguousLayout) Size() Size {
	return Size{c.Rows, c.Cols}
}

func (c ContiguousLayout) MatToVec(i, j int) int {
	if c.RowMajor {
		return i*c.Cols + j
	}
	return j*c.Rows + i
}

// Default (column-major) matricization.
func Mat(x vec.Const, rows, cols int) MatExpr {
	return MatColMajor(x, rows, cols)
}

// Addresses the elements of a vector as a matrix in column-major order.
func MatColMajor(x vec.Const, rows, cols int) MatExpr {
	return MatExpr{x, ContiguousLayout{rows, cols, false}}
}

// Addresses the elements of a vector as a matrix in column-major order.
func MatRowMajor(x vec.Const, rows, cols int) MatExpr {
	return MatExpr{x, ContiguousLayout{rows, cols, true}}
}

// Matricization. Upgrades a constant vector to have a size.
type MatExpr struct {
	Vector vec.Const
	Layout ContiguousLayout
}

func (expr MatExpr) Size() {
	return expr.Layout.Size()
}

func (expr MatExpr) At(i, j int) float64 {
	return expr.Vector.At(expr.Layout.MatToVec(i, j))
}

// Matricization. Upgrades a mutable vector to have a size.
type MutableMatExpr struct {
	Vector vec.Mutable
	Layout ContiguousLayout
}

func (expr MutableMatExpr) Size() {
	return expr.Layout.Size()
}

func (expr MutableMatExpr) At(i, j int) float64 {
	return expr.Vector.At(expr.Layout.MatToVec(i, j))
}

func (expr MutableMatExpr) Set(i, j int, x float64) {
	return expr.Vector.Set(expr.Layout.MatToVec(i, j), x)
}

// Lazily adds two matrices using default vectorization.
//
// Provided as a convenience to allow
//	C := mat.DenseCopy(mat.Plus(A, B))
// instead of
//	C := mat.MakeDense(Rows(A), Cols(A))
//	vec.Copy(mat.MutableVec(C), vec.Plus(mat.Vec(A), mat.Vec(B)))
func Plus(A, B Const) MatExpr {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return MatExpr{ContiguousLayout{rows, cols, false}, vec.Plus(Vec(A), Vec(B))}
}

// Lazily subtracts one matrix from another using default vectorization.
func Minus(A, B Const) MatExpr {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return MatExpr{ContiguousLayout{rows, cols, false}, vec.Minus(Vec(A), Vec(B))}
}

// Lazily scales a matrix using default vectorization.
func Scale(k float64, A Const) MatExpr {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return MatExpr{ContiguousLayout{rows, cols, false}, vec.Scale(k, Vec(A))}
}

// Lazy element-wise multiplication of two matrices using default vectorization.
func VectorMultiply(A, B Const) MatExpr {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	rows, cols := RowsCols(A)
	return MatExpr{ContiguousLayout{rows, cols, false}, vec.Times(Vec(A), Vec(B))}
}
