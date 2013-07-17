package vec

// Expressions which don't have a Mutable partner.

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

// Compute x^2 for every element in the vector.
// Lazily evaluated.
func Square(x Const) Const {
	return Multiply(x, x)
}

// Addition of two vectors.
// Lazily evaluated.
func Plus(xs ...Const) Const {
	return MapN(Sum, xs...)
}

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) Const {
	f := func(x, y float64) float64 { return x - y }
	return MapTwo(x, y, f)
}

// Element-wise multiplication.
// Lazily evaluated.
func Multiply(xs ...Const) Const {
	return MapN(Prod, xs...)
}

// Element-wise division.
// Lazily evaluated.
func Divide(x, y Const) Const {
	f := func(x, y float64) float64 { return x / y }
	return MapTwo(x, y, f)
}

// Constant vector of ones.
func Ones(n int) Const {
	return Constant(n, 1)
}

// Constant vector of zeros.
func Zeros(n int) Const {
	return Constant(n, 0)
}

// Constant vector.
func Constant(n int, a float64) Const {
	f := func() float64 { return a }
	return MapNil(n, f)
}
