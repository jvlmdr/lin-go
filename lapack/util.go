package lapack

import "github.com/jackvalmadre/lin-go/mat"

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func re(x interface{}) float64 {
	switch x := x.(type) {
	default:
	case float64:
		return x
	case complex128:
		return real(x)
	}
	panic("neither float64 not complex128")
}

func appendRows(A mat.Stride, n int) mat.Stride {
	rows := A.Rows + n
	// Re-allocate if more space required.
	if !A.InCapTo(rows, A.Cols) {
		rowcap := max(2*A.Stride, A.Rows)
		tmp := mat.MakeStrideCap(A.Rows, A.Cols, rowcap, A.Cols)
		mat.Copy(tmp, A)
		A = tmp
	}
	A = A.SliceTo(rows, A.Cols)
	// Ensure zero.
	for i := A.Rows; i < rows; i++ {
		for j := 0; j < A.Cols; j++ {
			A.Set(i, j, 0)
		}
	}
	return A
}
