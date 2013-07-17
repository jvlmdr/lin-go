package zvec

// This file contains expressions which manipulate the shape of vectors.

// Concatenates vectors.
func Cat(x ...Const) Const {
	if len(x) == 0 {
		return nil
	}
	if len(x) == 1 {
		return x[0]
	}
	if len(x) == 2 {
		return CatTwo(x[0], x[1])
	}
	// Split at midpoint.
	m := (len(x) + 1) / 2
	return CatTwo(Cat(x[:m]...), Cat(x[m:]...))
}

// Concatenates two vectors.
func CatTwo(x, y Const) Const {
	return catExpr{x, y, x.Len() + y.Len()}
}

// It is important that this structure caches its length to achieve O(log N) lookup.
type catExpr struct {
	X Const
	Y Const
	N int
}

func (expr catExpr) Len() int { return expr.N }

func (expr catExpr) At(i int) complex128 {
	m := expr.X.Len()
	if i < m {
		return expr.X.At(i)
	}
	return expr.Y.At(i - m)
}

// Concatenates vectors.
func MutableCat(x ...Mutable) Mutable {
	if len(x) == 0 {
		return nil
	}
	if len(x) == 1 {
		return x[0]
	}
	if len(x) == 2 {
		return MutableCatTwo(x[0], x[1])
	}
	// Split at midpoint.
	m := (len(x) + 1) / 2
	return MutableCatTwo(MutableCat(x[:m]...), MutableCat(x[m:]...))
}

// Concatenates two vectors.
func MutableCatTwo(x, y Mutable) Mutable {
	return mutableCatExpr{x, y, x.Len() + y.Len()}
}

type mutableCatExpr struct {
	X Mutable
	Y Mutable
	N int
}

func (expr mutableCatExpr) Len() int { return expr.N }

func (expr mutableCatExpr) At(i int) complex128 {
	m := expr.X.Len()
	if i < m {
		return expr.X.At(i)
	}
	return expr.Y.At(i - m)
}

func (expr mutableCatExpr) Set(i int, x complex128) {
	m := expr.X.Len()
	if i < m {
		expr.X.Set(i, x)
		return
	}
	expr.Y.Set(i-m, x)
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
