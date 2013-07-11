package vec

import "math/rand"

// Lazy evaluators for vector-valued functions.
// Idiomatic use is x <- x - y written vec.Copy(x, vec.Minus(x, y)).

// Addition of two vectors.
// Lazily evaluated.
func Plus(x, y Const) Const { return plusExpr{x, y} }

type plusExpr struct{ X, Y Const }

func (expr plusExpr) Size() int        { return expr.X.Size() }
func (expr plusExpr) At(i int) float64 { return expr.X.At(i) + expr.Y.At(i) }

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) Const { return minusExpr{x, y} }

type minusExpr struct{ X, Y Const }

func (expr minusExpr) Size() int        { return expr.X.Size() }
func (expr minusExpr) At(i int) float64 { return expr.X.At(i) - expr.Y.At(i) }

// Multiplication by a scalar.
// Lazily evaluated.
func Scale(alpha float64, x Const) Const { return scaleExpr{alpha, x} }

type scaleExpr struct {
	Alpha float64
	X     Const
}

func (expr scaleExpr) Size() int        { return expr.X.Size() }
func (expr scaleExpr) At(i int) float64 { return expr.Alpha * expr.X.At(i) }

// Element-wise multiplication.
// Lazily evaluated.
func Multiply(x, y Const) Const { return multiplyExpr{x, y} }

type multiplyExpr struct{ X, Y Const }

func (expr multiplyExpr) Size() int        { return expr.X.Size() }
func (expr multiplyExpr) At(i int) float64 { return expr.X.At(i) * expr.Y.At(i) }

// Constant vector of ones.
func One(n int) Const { return oneExpr(n) }

type oneExpr int

func (expr oneExpr) Size() int      { return int(expr) }
func (expr oneExpr) At(int) float64 { return 1 }

// Constant vector of zeros.
func Zero(n int) Const { return zeroExpr(n) }

type zeroExpr int

func (expr zeroExpr) Size() int      { return int(expr) }
func (expr zeroExpr) At(int) float64 { return 0 }

// Constant vector.
func Constant(n int, x float64) Const { return constantExpr{n, x} }

type constantExpr struct {
	N int
	X float64
}

func (expr constantExpr) Size() int      { return expr.N }
func (expr constantExpr) At(int) float64 { return expr.X }

// Applies the same unary function to every element.
// Lazily evaluated.
func Map(x Const, f func(int, float64) float64) Const { return mapExpr{x, f} }

type mapExpr struct {
	X Const
	F func(int, float64) float64
}

func (expr mapExpr) Size() int        { return expr.X.Size() }
func (expr mapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i)) }

// Applies the same binary function to every element.
// Lazily evaluated.
func BinaryMap(x, y Const, f func(int, float64, float64) float64) Const {
	return binaryMapExpr{x, y, f}
}

type binaryMapExpr struct {
	X Const
	Y Const
	F func(int, float64, float64) float64
}

func (expr binaryMapExpr) Size() int        { return expr.X.Size() }
func (expr binaryMapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i), expr.Y.At(i)) }

// Lazily concatenates two vectors.
func Cat(x, y Const) Const {
	return catExpr{x, y}
}

type catExpr struct{ X, Y Const }

func (expr catExpr) Size() int { return expr.X.Size() + expr.Y.Size() }

func (expr catExpr) At(i int) float64 {
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

func (expr mutableCatExpr) At(i int) float64 {
	n := expr.X.Size()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

func (expr mutableCatExpr) Set(i int, x float64) {
	n := expr.X.Size()
	if i < n {
		expr.X.Set(i, x)
		return
	}
	expr.Y.Set(i-n, x)
}

// Vector whose entries are random and normally distributed.
func Randn(n int) Const { return randnExpr(n) }

type randnExpr int

func (expr randnExpr) Size() int      { return int(expr) }
func (expr randnExpr) At(int) float64 { return rand.NormFloat64() }
