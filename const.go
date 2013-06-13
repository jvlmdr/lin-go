package vec

import "math"

func Dot(x, y Const) float64 {
	var acc float64
	for i := 0; i < x.Size(); i++ {
		acc += x.At(i) * y.At(i)
	}
	return acc
}

func SqrNorm(x Const) float64 {
	return Dot(x, x)
}

func Norm(x Const) float64 {
	return math.Sqrt(SqrNorm(x))
}

func InfNorm(x Const) float64 {
	var acc float64
	for i := 0; i < x.Size(); i++ {
		acc = math.Max(acc, math.Abs(x.At(i)))
	}
	return acc
}

func SquareDistance(x, y Const) float64 {
	var acc float64
	for i := 0; i < x.Size(); i++ {
		d := x.At(i) - y.At(i)
		acc += d * d
	}
	return acc
}

func Distance(x, y Const) float64 {
	return math.Sqrt(SquareDistance(x, y))
}
