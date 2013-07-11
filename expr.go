package vec

import "math/rand"

// Lazy evaluators for vector-valued functions.
// Idiomatic use is x <- x - y written vec.Copy(x, vec.Minus(x, y)).

// Addition of two vectors.
// Lazily evaluated.
func Plus(x, y Const) PlusExpr { return PlusExpr{x, y} }

type PlusExpr struct{ X, Y Const }

func (expr PlusExpr) Size() int        { return expr.X.Size() }
func (expr PlusExpr) At(i int) float64 { return expr.X.At(i) + expr.Y.At(i) }

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) MinusExpr { return MinusExpr{x, y} }

type MinusExpr struct{ X, Y Const }

func (expr MinusExpr) Size() int        { return expr.X.Size() }
func (expr MinusExpr) At(i int) float64 { return expr.X.At(i) - expr.Y.At(i) }

// Multiplication by a scalar.
// Lazily evaluated.
func Scale(alpha float64, x Const) ScaleExpr { return ScaleExpr{alpha, x} }

type ScaleExpr struct {
	Alpha float64
	X     Const
}

func (expr ScaleExpr) Size() int        { return expr.X.Size() }
func (expr ScaleExpr) At(i int) float64 { return expr.Alpha * expr.X.At(i) }

// Element-wise multiplication.
// Lazily evaluated.
func Multiply(x, y Const) MultiplyExpr { return MultiplyExpr{x, y} }

type MultiplyExpr struct{ X, Y Const }

func (expr MultiplyExpr) Size() int        { return expr.X.Size() }
func (expr MultiplyExpr) At(i int) float64 { return expr.X.At(i) * expr.Y.At(i) }

// Constant vector of ones.
func One(n int) OneExpr { return OneExpr{n} }

type OneExpr struct{ N int }

func (expr OneExpr) Size() int      { return expr.N }
func (expr OneExpr) At(int) float64 { return 1 }

// Constant vector of zeros.
func Zero(n int) ZeroExpr { return ZeroExpr{n} }

type ZeroExpr struct{ N int }

func (expr ZeroExpr) Size() int      { return expr.N }
func (expr ZeroExpr) At(int) float64 { return 0 }

// Constant vector.
func Constant(n int, x float64) ConstantExpr { return ConstantExpr{n, x} }

type ConstantExpr struct {
	N int
	X float64
}

func (expr ConstantExpr) Size() int      { return expr.N }
func (expr ConstantExpr) At(int) float64 { return expr.X }

// Applies the same unary function to every element.
// Lazily evaluated.
func Map(x Const, f func(int, float64) float64) MapExpr { return MapExpr{x, f} }

type MapExpr struct {
	X Const
	F func(int, float64) float64
}

func (expr MapExpr) Size() int        { return expr.X.Size() }
func (expr MapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i)) }

// Applies the same binary function to every element.
// Lazily evaluated.
func BinaryMap(x, y Const, f func(int, float64, float64) float64) BinaryMapExpr {
	return BinaryMapExpr{x, y, f}
}

type BinaryMapExpr struct {
	X Const
	Y Const
	F func(int, float64, float64) float64
}

func (expr BinaryMapExpr) Size() int        { return expr.X.Size() }
func (expr BinaryMapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i), expr.Y.At(i)) }

// Lazily concatenates two vectors.
func Cat(x, y Const) CatExpr {
	return CatExpr{x, y}
}

type CatExpr struct{ X, Y Const }

func (expr CatExpr) Size() int { return expr.X.Size() + expr.Y.Size() }

func (expr CatExpr) At(i int) float64 {
	n := expr.X.Size()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

// Lazily concatenates two mutable vectors.
func MutableCat(x, y Mutable) MutableCatExpr {
	return MutableCatExpr{x, y}
}

type MutableCatExpr struct{ X, Y Mutable }

func (expr MutableCatExpr) Size() int { return expr.X.Size() + expr.Y.Size() }

func (expr MutableCatExpr) At(i int) float64 {
	n := expr.X.Size()
	if i < n {
		return expr.X.At(i)
	}
	return expr.Y.At(i - n)
}

func (expr MutableCatExpr) Set(i int, x float64) {
	n := expr.X.Size()
	if i < n {
		expr.X.Set(i, x)
		return
	}
	expr.Y.Set(i-n, x)
}

// Vector whose entries are random and normally distributed.
func Randn(n int) RandnExpr { return RandnExpr(n) }

type RandnExpr int

func (expr RandnExpr) Size() int { return int(expr) }

func (expr RandnExpr) At(int) float64 {
	return rand.NormFloat64()
}
