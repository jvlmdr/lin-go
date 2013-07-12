package zmat

// This file contains things you can do with a Const matrix.

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
