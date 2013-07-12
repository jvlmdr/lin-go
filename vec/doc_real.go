/*
Package vec provides common functionality for working with vectors.

The core of the package is the vec.Const interface, which has Size() and At(i, j) methods, and the vec.Mutable interface, which adds Set(i, j, x).
These interfaces combined with a Copy(Mutable, Const) method and thin wrappers for Plus(), Minus(), etc. provide idiomatic use:
	c := vec.Make(n)
	vec.Copy(c, vec.Plus(a, b))

	// Succinct version:
	c := vec.MakeCopy(vec.Plus(a, b))

	// In-place version:
	vec.Copy(a, vec.Plus(a, b))
These thin wrappers also give a neat mechanism for chaining operations:
	// Gradient descent update:
	vec.Copy(x, vec.Minus(x, vec.Scale(alpha, dfdx)))

	// Compute root-mean-square magnitude.
	rms := math.Sqrt(vec.Mean(vec.Square(x)))
The package contains a single concrete implementation of a vector, vec.Slice, which simply augments a []float64 with the necessary methods.
*/
package vec
