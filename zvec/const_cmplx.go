package zvec

// Miscellaneous operations you can do with a Const vector.
// Operations which are not easily converted from real to complex by sed.

import (
	"github.com/jackvalmadre/lin-go/vec"
	"math"
)

func Mean(x Const) complex128 {
	return Sum(x) * complex(1/float64(x.Len()), 0)
}

// Computes x^H y = sum_i conj(x_i) y_i.
// Note that the complex dot product is not commutative.
// Dot(x, y) = cmplx.Conj(Dot(y, x))
func Dot(x, y Const) complex128 {
	return Sum(Multiply(Conj(x), y))
}

func SqrNorm(x Const) float64 {
	return vec.Sum(vec.Square(Abs(x)))
}

func Norm(x Const) float64 {
	return math.Sqrt(SqrNorm(x))
}

func SqrDist(x, y Const) float64 {
	return SqrNorm(Minus(x, y))
}

func Dist(x, y Const) float64 {
	return math.Sqrt(SqrDist(x, y))
}

func OneNorm(x Const) float64 {
	return vec.Sum(Abs(x))
}

func InfNorm(x Const) float64 {
	y, _ := vec.Max(Abs(x))
	return y
}

func LpNorm(x Const, p float64) float64 {
	return math.Pow(vec.Sum(vec.Pow(Abs(x), p)), 1/p)
}
