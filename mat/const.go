package mat

// This file contains things you can do with a Const matrix.

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Returns the number of rows in A.
func Rows(A Const) int {
	return A.Size().Rows
}

// Returns the number of columns in A.
func Cols(A Const) int {
	return A.Size().Cols
}

// Returns the number of rows and columns in A.
func RowsCols(A Const) (int, int) {
	s := A.Size()
	return s.Rows, s.Cols
}

// Returns the bounds of the matrix as a rectangle starting at (0, 0).
func Bounds(A Const) Rect {
	rows, cols := RowsCols(A)
	return MakeRect(0, 0, rows, cols)
}

// Prints all elements using format (e.g. "%g").
func Fprintf(w io.Writer, format string, A Const) {
	m, n := RowsCols(A)
	for i := 0; i < m; i++ {
		if i > 0 {
			fmt.Fprint(w, "\n")
		}
		for j := 0; j < n; j++ {
			fmt.Fprintf(w, format, A.At(i, j))
		}
	}
}

// Returns Sprintf(A, "%g").
func String(A Const) string {
	return Sprintf("%g", A)
}

// See Fprintf.
func Sprintf(format string, A Const) string {
	var b bytes.Buffer
	Fprintf(&b, format, A)
	return b.String()
}

// See Fprintf.
func Printf(format string, A Const) {
	Fprintf(os.Stdout, format, A)
}
