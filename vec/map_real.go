package vec

import (
	"math"
	"math/rand"
)

// Compute |x| for every element in the vector.
// Lazily evaluated.
func Abs(x Const) Const {
	return Map(x, math.Abs)
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

// Vector whose entries are random and normally distributed.
func Randn(n int) Const {
	f := func() float64 { return rand.NormFloat64() }
	return MapNil(n, f)
}
