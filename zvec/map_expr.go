package zvec

// This file constains constant vector expressions expressed as a map.

// Vector whose i-th element is f(x.At(i)).
func Map(x Const, f func(complex128) complex128) Const { return mapExpr{x, f} }

type mapExpr struct {
	X Const
	F func(complex128) complex128
}

func (expr mapExpr) Len() int            { return expr.X.Len() }
func (expr mapExpr) At(i int) complex128 { return expr.F(expr.X.At(i)) }

// Vector whose i-th element is f(x.At(i), y.At(i)).
func MapTwo(x, y Const, f func(complex128, complex128) complex128) Const {
	return mapTwoExpr{x, y, f}
}

type mapTwoExpr struct {
	X Const
	Y Const
	F func(complex128, complex128) complex128
}

func (expr mapTwoExpr) Len() int            { return expr.X.Len() }
func (expr mapTwoExpr) At(i int) complex128 { return expr.F(expr.X.At(i), expr.Y.At(i)) }

// Vector whose i-th element is f(vec.Slice([]complex128{x[0].At(i), ..., x[n-1].At(i)}))
func MapN(f func(Const) complex128, xs ...Const) Const {
	if len(xs) == 0 {
		panic("Empty list of vectors")
	}
	panicIfNotSameLength(xs...)
	return mapNExpr{xs, f}
}

type mapNExpr struct {
	X []Const
	F func(Const) complex128
}

func (expr mapNExpr) Len() int            { return expr.X[0].Len() }
func (expr mapNExpr) At(i int) complex128 { return expr.F(sameElementMultipleVectors{expr.X, i}) }

type sameElementMultipleVectors struct {
	X []Const
	I int
}

func (expr sameElementMultipleVectors) Len() int            { return len(expr.X) }
func (expr sameElementMultipleVectors) At(i int) complex128 { return expr.X[i].At(expr.I) }

// Vector whose i-th element is f().
func MapNil(n int, f func() complex128) Const { return mapNilExpr{n, f} }

type mapNilExpr struct {
	N int
	F func() complex128
}

func (expr mapNilExpr) Len() int            { return expr.N }
func (expr mapNilExpr) At(i int) complex128 { return expr.F() }

// Vector whose i-th element is f(i).
func MapIndex(n int, f func(int) complex128) Const { return mapIndexExpr{n, f} }

type mapIndexExpr struct {
	N int
	F func(int) complex128
}

func (expr mapIndexExpr) Len() int            { return expr.N }
func (expr mapIndexExpr) At(i int) complex128 { return expr.F(i) }

// Vector whose i-th element is f(i) and modified by g(i, x).
func MutableMapIndex(n int, f func(int) complex128, g func(int, complex128)) Mutable {
	return mutableMapIndexExpr{n, f, g}
}

type mutableMapIndexExpr struct {
	N int
	F func(int) complex128
	G func(int, complex128)
}

func (expr mutableMapIndexExpr) Len() int                { return expr.N }
func (expr mutableMapIndexExpr) At(i int) complex128     { return expr.F(i) }
func (expr mutableMapIndexExpr) Set(i int, x complex128) { expr.G(i, x) }
