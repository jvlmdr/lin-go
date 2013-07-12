package vec

// Miscellaneous operations you can do with a Const vector.

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Sum(x Const) float64 {
	var total float64
	for i := 0; i < x.Size(); i++ {
		total += x.At(i)
	}
	return total
}

func Append(s []float64, x Const) []float64 {
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
