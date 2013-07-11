package vec

// Miscellaneous operations you can do with a Const vector.

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
)

func Sum(x Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		total += x.At(i)
	}
	return total
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

func Min(x Const) (float64, int) {
	argmin := -1
	xmin := math.Inf(1)
	for i := 0; i < x.Size(); i++ {
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
	for i := 0; i < x.Size(); i++ {
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
	for i := 0; i < x.Size(); i++ {
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

func InfNorm(x Const) float64 {
	y, _ := Max(Abs(x))
	return y
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
