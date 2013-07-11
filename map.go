package vec

// This file contains element-wise operations.

import (
	"math"
	"math/rand"
)

// Lazy evaluators for vector-valued functions.
// Idiomatic use is x <- x - y written vec.Copy(x, vec.Minus(x, y)).

// Multiplication by a scalar.
// Lazily evaluated.
func Scale(a float64, x Const) Const {
	f := func(x float64) float64 { return a * x }
	return Map(x, f)
}

// Compute 1/x for every element in the vector.
// Lazily evaluated.
func Invert(x Const) Const {
	return Ldivide(x, 1)
}

// Compute a/x for every element in the vector.
// Lazily evaluated.
func Ldivide(x Const, a float64) Const {
	f := func(x float64) float64 { return a / x }
	return Map(x, f)
}

// Compute |x| for every element in the vector.
// Lazily evaluated.
func Abs(x Const) Const {
	return Map(x, math.Abs)
}

// Compute x^2 for every element in the vector.
// Lazily evaluated.
func Sqr(x Const) Const {
	f := func(x float64) float64 { return x * x }
	return Map(x, f)
}

// Compute exp(x) for every element in the vector.
// Lazily evaluated.
func Exp(x Const) Const {
	return Map(x, math.Exp)
}

// Compute log(x) for every element in the vector.
// Lazily evaluated.
func Log(x Const) Const {
	return Map(x, math.Log)
}

// Compute sqrt(x) for every element in the vector.
// Lazily evaluated.
func Sqrt(x Const) Const {
	return Map(x, math.Sqrt)
}

// Compute x^p for every element in the vector.
// Lazily evaluated.
func Pow(x Const, p float64) Const {
	f := func(x float64) float64 { return math.Pow(x, p) }
	return Map(x, f)
}

// Addition of two vectors.
// Lazily evaluated.
func Plus(x, y Const) Const {
	f := func(x, y float64) float64 { return x + y }
	return BinaryMap(x, y, f)
}

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) Const {
	f := func(x, y float64) float64 { return x - y }
	return BinaryMap(x, y, f)
}

// Element-wise multiplication.
// Lazily evaluated.
func Multiply(x, y Const) Const {
	f := func(x, y float64) float64 { return x * y }
	return BinaryMap(x, y, f)
}

// Element-wise division.
// Lazily evaluated.
func Divide(x, y Const) Const {
	f := func(x, y float64) float64 { return x / y }
	return BinaryMap(x, y, f)
}

// Constant vector of ones.
func One(n int) Const {
	return Constant(n, 1)
}

// Constant vector of zeros.
func Zero(n int) Const {
	return Constant(n, 0)
}

// Constant vector.
func Constant(n int, a float64) Const {
	f := func() float64 { return a }
	return IndexMap(n, f)
}

// Vector whose entries are random and normally distributed.
func Randn(n int) Const {
	f := func() float64 { return rand.NormFloat64() }
	return IndexMap(n, f)
}

// Applies the same zero-ary function for every argument.
// Lazily evaluated.
func IndexMap(n int, f func() float64) Const { return indexMapExpr{n, f} }

type indexMapExpr struct {
	N int
	F func() float64
}

func (expr indexMapExpr) Size() int        { return expr.N }
func (expr indexMapExpr) At(i int) float64 { return expr.F() }

// Applies the same unary function to every element.
// Lazily evaluated.
func Map(x Const, f func(float64) float64) Const { return mapExpr{x, f} }

type mapExpr struct {
	X Const
	F func(float64) float64
}

func (expr mapExpr) Size() int        { return expr.X.Size() }
func (expr mapExpr) At(i int) float64 { return expr.F(expr.X.At(i)) }

// Applies the same binary function to every element.
// Lazily evaluated.
func BinaryMap(x, y Const, f func(float64, float64) float64) Const {
	return binaryMapExpr{x, y, f}
}

type binaryMapExpr struct {
	X Const
	Y Const
	F func(float64, float64) float64
}

func (expr binaryMapExpr) Size() int        { return expr.X.Size() }
func (expr binaryMapExpr) At(i int) float64 { return expr.F(expr.X.At(i), expr.Y.At(i)) }
