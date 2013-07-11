package zvec

// This file contains element-wise operations.

// Lazy evaluators for vector-valued functions.
// Idiomatic use is x <- x - y written vec.Copy(x, vec.Minus(x, y)).

// Multiplication by a scalar.
// Lazily evaluated.
func Scale(a complex128, x Const) Const {
	f := func(x complex128) complex128 { return a * x }
	return Map(x, f)
}

// Compute 1/x for every element in the vector.
// Lazily evaluated.
func Invert(x Const) Const {
	return Ldivide(x, 1)
}

// Compute a/x for every element in the vector.
// Lazily evaluated.
func Ldivide(x Const, a complex128) Const {
	f := func(x complex128) complex128 { return a / x }
	return Map(x, f)
}

// Compute x^2 for every element in the vector.
// Lazily evaluated.
func Square(x Const) Const {
	f := func(x complex128) complex128 { return x * x }
	return Map(x, f)
}

// Addition of two vectors.
// Lazily evaluated.
func Plus(x, y Const) Const {
	f := func(x, y complex128) complex128 { return x + y }
	return BinaryMap(x, y, f)
}

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) Const {
	f := func(x, y complex128) complex128 { return x - y }
	return BinaryMap(x, y, f)
}

// Element-wise multiplication.
// Lazily evaluated.
func Multiply(x, y Const) Const {
	f := func(x, y complex128) complex128 { return x * y }
	return BinaryMap(x, y, f)
}

// Element-wise division.
// Lazily evaluated.
func Divide(x, y Const) Const {
	f := func(x, y complex128) complex128 { return x / y }
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
func Constant(n int, a complex128) Const {
	f := func() complex128 { return a }
	return ConstantMap(n, f)
}

// Applies the same zero-ary function for every argument.
// Lazily evaluated.
func ConstantMap(n int, f func() complex128) Const { return constantMapExpr{n, f} }

type constantMapExpr struct {
	N int
	F func() complex128
}

func (expr constantMapExpr) Size() int           { return expr.N }
func (expr constantMapExpr) At(i int) complex128 { return expr.F() }

// Applies the same unary function to every element.
// Lazily evaluated.
func Map(x Const, f func(complex128) complex128) Const { return mapExpr{x, f} }

type mapExpr struct {
	X Const
	F func(complex128) complex128
}

func (expr mapExpr) Size() int           { return expr.X.Size() }
func (expr mapExpr) At(i int) complex128 { return expr.F(expr.X.At(i)) }

// Applies the same binary function to every element.
// Lazily evaluated.
func BinaryMap(x, y Const, f func(complex128, complex128) complex128) Const {
	return binaryMapExpr{x, y, f}
}

type binaryMapExpr struct {
	X Const
	Y Const
	F func(complex128, complex128) complex128
}

func (expr binaryMapExpr) Size() int           { return expr.X.Size() }
func (expr binaryMapExpr) At(i int) complex128 { return expr.F(expr.X.At(i), expr.Y.At(i)) }
