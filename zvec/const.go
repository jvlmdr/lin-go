package zvec

// Miscellaneous operations you can do with a Const vector.

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Sum(x Const) complex128 {
	var total complex128
	for i := 0; i < x.Len(); i++ {
		total += x.At(i)
	}
	return total
}

func Append(s []complex128, x Const) []complex128 {
	n := len(s) + x.Len()
	// Re-allocate only once if at all.
	if n > cap(s) {
		// At least double the previous capacity.
		t := make([]complex128, len(s), max(n, 2*cap(s)))
		copy(t, s)
		s = t
	}
	// Append new elements.
	for i := 0; i < x.Len(); i++ {
		s = append(s, x.At(i))
	}
	return s
}

// Prints all elements using format (e.g. "%g").
// Result is "(x.At(0), x.At(1), ..., x.At(n-1))".
func Fprintf(w io.Writer, format string, x Const) {
	fmt.Fprintf(w, "(")
	for i := 0; i < x.Len(); i++ {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, format, x.At(i))
	}
	fmt.Fprintf(w, ")")
}

// Returns Sprintf("%g", x).
func String(x Const) string {
	return Sprintf("%g", x)
}

// See Fprintf.
func Sprintf(format string, x Const) string {
	var b bytes.Buffer
	Fprintf(&b, format, x)
	return b.String()
}

// See Fprintf.
func Printf(format string, x Const) {
	Fprintf(os.Stdout, format, x)
}
