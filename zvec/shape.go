package zvec

// This file contains operations involving the shape of a vector.

// Lazily concatenates two vectors.
func Cat(x, y Const) Const {
	return catExpr{x, y}
}

type catExpr struct{ X, Y Const }

func (expr catExpr) Size() int { return expr.X.Size() + expr.Y.Size() }

func (expr catExpr) At(i int) complex128 {
	n := expr.X.Size()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

// Lazily concatenates two mutable vectors.
func MutableCat(x, y Mutable) Mutable {
	return mutableCatExpr{x, y}
}

type mutableCatExpr struct{ X, Y Mutable }

func (expr mutableCatExpr) Size() int { return expr.X.Size() + expr.Y.Size() }

func (expr mutableCatExpr) At(i int) complex128 {
	n := expr.X.Size()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

func (expr mutableCatExpr) Set(i int, x complex128) {
	n := expr.X.Size()
	if i < n {
		expr.X.Set(i, x)
		return
	}
	expr.Y.Set(i-n, x)
}
