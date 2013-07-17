package zvec

// This file contains expressions which manipulate the shape of vectors.

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
	at := func(i int) complex128 {
		return x.At(i - a)
	}
	return MapIndex(b-a, at)
}

// Subvector of elements in [a, b).
func MutableSubvec(x Mutable, a, b int) Mutable {
	at := func(i int) complex128 {
		return x.At(i - a)
	}
	set := func(i int, v complex128) {
		x.Set(i-a, v)
	}
	return MutableMapIndex(b-a, at, set)
}
