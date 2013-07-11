package vec

// Miscellaneous operations you can do with a Const vector.

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
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

func SqrDist(x, y Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		d := x.At(i) - y.At(i)
		total += d * d
	}
	return total
}

func Dist(x, y Const) float64 {
	return math.Sqrt(SqrDist(x, y))
}

func AppendToSlice(s []float64, x Const) []float64 {
	n := len(s) + x.Size()
	// Re-allocate only once if at all.
	if n > cap(s) {
		// At least double the previous capacity.
		t := make([]float64, len(s), max(n, 2*cap(s)))
		copy(t, s)
		s = t
	}
	// Append new elements.
	for i := 0; i < x.Size(); i++ {
		s = append(s, x.At(i))
	}
	return s
}

func Fprint(w io.Writer, x Const, format string) {
	fmt.Fprint(w, "(")
	for i := 0; i < x.Size(); i++ {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, format, x.At(i))
	}
	fmt.Fprint(w, ")")
}

func String(x Const) string {
	var b bytes.Buffer
	Fprint(&b, x, "%g")
	return b.String()
}

func Print(x Const, format string) {
	Fprint(os.Stdout, x, format)
}
