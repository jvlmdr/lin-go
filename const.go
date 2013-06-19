package vec

import (
	"bytes"
	"fmt"
	"math"
)

func Dot(x, y Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		total += x.At(i) * y.At(i)
	}
	return total
}

func SqrNorm(x Const) float64 {
	return Dot(x, x)
}

func Norm(x Const) float64 {
	return math.Sqrt(SqrNorm(x))
}

func Sum(x Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		total += x.At(i)
	}
	return total
}

func OneNorm(x Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		total += math.Abs(x.At(i))
	}
	return total
}

func Min(x Const) float64 {
	min := math.Inf(1)
	for i := 0; i < x.Size(); i++ {
		min = math.Min(min, x.At(i))
	}
	return min
}

func Max(x Const) float64 {
	max := math.Inf(-1)
	for i := 0; i < x.Size(); i++ {
		max = math.Max(max, x.At(i))
	}
	return max
}

func MinMax(x Const) (float64, float64) {
	min := math.Inf(1)
	max := math.Inf(-1)
	for i := 0; i < x.Size(); i++ {
		min = math.Min(min, x.At(i))
		max = math.Max(max, x.At(i))
	}
	return min, max
}

func InfNorm(x Const) float64 {
	var max float64
	for i := 0; i < x.Size(); i++ {
		max = math.Max(max, math.Abs(x.At(i)))
	}
	return max
}

func SquareDistance(x, y Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		d := x.At(i) - y.At(i)
		total += d * d
	}
	return total
}

func Distance(x, y Const) float64 {
	return math.Sqrt(SquareDistance(x, y))
}

func String(x Const) string {
	var b bytes.Buffer
	b.WriteString("(")
	for i := 0; i < x.Size(); i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%g", x.At(i))
	}
	b.WriteString(")")
	return b.String()
}
