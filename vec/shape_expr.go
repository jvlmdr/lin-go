package vec

// This file contains expressions which manipulate the shape of vectors.

// Lazily concatenates two vectors.
func Cat(x, y Const) Const {
	return catExpr{x, y}
}

type catExpr struct{ X, Y Const }

func (expr catExpr) Len() int { return expr.X.Len() + expr.Y.Len() }

func (expr catExpr) At(i int) float64 {
	n := expr.X.Len()
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

func (expr mutableCatExpr) Len() int { return expr.X.Len() + expr.Y.Len() }

func (expr mutableCatExpr) At(i int) float64 {
	n := expr.X.Len()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

func (expr mutableCatExpr) Set(i int, x float64) {
	n := expr.X.Len()
	if i < n {
		expr.X.Set(i, x)
		return
	}
	expr.Y.Set(i-n, x)
}

// Concatenate many vectors.
func CatN(x ...Const) Const {
	if len(x) == 0 {
		return nil
	}
	if len(x) == 1 {
		return x[0]
	}
	m := (len(x) + 1) / 2
	return Cat(CatN(x[:m]...), CatN(x[m:]...))
}

// Subvector of elements in [a, b).
func Subvec(x Const, a, b int) Const {
	at := func(i int) float64 {
		return x.At(i - a)
	}
	return MapIndex(b-a, at)
}

// Subvector of elements in [a, b).
func MutableSubvec(x Mutable, a, b int) Mutable {
	at := func(i int) float64 {
		return x.At(i - a)
	}
	set := func(i int, v float64) {
		x.Set(i-a, v)
	}
	return MutableMapIndex(b-a, at, set)
}
