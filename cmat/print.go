package cmat

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Prints all elements using format (e.g. "%g").
func Fprintf(w io.Writer, format string, a Const) {
	m, n := a.Dims()
	for i := 0; i < m; i++ {
		if i > 0 {
			fmt.Fprint(w, "\n")
		}
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, format, a.At(i, j))
		}
	}
}

// Returns Sprintf(a, "%g").
func String(a Const) string {
	return Sprintf("%g", a)
}

// See Fprintf.
func Sprintf(format string, a Const) string {
	var b bytes.Buffer
	Fprintf(&b, format, a)
	return b.String()
}

// See Fprintf.
func Printf(format string, a Const) {
	Fprintf(os.Stdout, format, a)
}
