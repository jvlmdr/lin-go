package vec

// Miscellaneous operations you can do with a Const vector.

import "math"

func Mean(x Const) float64 {
	return Sum(x) / float64(x.Len())
}

func Dot(x, y Const) float64 {
	return Sum(Multiply(x, y))
}

func SqrNorm(x Const) float64 {
	return Dot(x, x)
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
	return Sum(Abs(x))
}

func InfNorm(x Const) float64 {
	y, _ := Max(Abs(x))
	return y
}

func LpNorm(x Const, p float64) float64 {
	return math.Pow(Sum(Pow(Abs(x), p)), 1/p)
}

func Min(x Const) (float64, int) {
	argmin := -1
	xmin := math.Inf(1)
	for i := 0; i < x.Len(); i++ {
		xi := x.At(i)
		if xi < xmin {
			xmin = xi
			argmin = i
		}
	}
	return xmin, argmin
}

func Max(x Const) (float64, int) {
	argmax := -1
	xmax := math.Inf(-1)
	for i := 0; i < x.Len(); i++ {
		xi := x.At(i)
		if xi > xmax {
			xmax = xi
			argmax = i
		}
	}
	return xmax, argmax
}

func MinMax(x Const) (xmin float64, argmin int, xmax float64, argmax int) {
	argmin = -1
	argmax = -1
	xmin = math.Inf(1)
	xmax = math.Inf(-1)
	for i := 0; i < x.Len(); i++ {
		xi := x.At(i)
		if xi < xmin {
			xmin = xi
			argmin = i
		}
		if xi > xmax {
			xmax = xi
			argmax = i
		}
	}
	return
}
