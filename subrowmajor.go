package mat

import "github.com/jackvalmadre/go-vec"

// Submatrix within a contiguous matrix, column-major order.
type SubContiguousRowMajor struct {
	Rows   int
	Cols   int
	Stride int
	// The (i, j)-th element is at Elements[i*Stride+j].
	Elements []float64
}

func (A SubContiguousRowMajor) Size() Size {
	return Size{A.Rows, A.Cols}
}

func (A SubContiguousRowMajor) At(i, j int) float64 {
	return A.Elements[i*A.Stride+j]
}

func (A SubContiguousRowMajor) Set(i, j int, x float64) {
	A.Elements[i*A.Stride+j] = x
}

func (A SubContiguousRowMajor) RowMajor() bool {
	return true
}

func (A SubContiguousRowMajor) Array() []float64 {
	return A.Elements
}

// Returns the transpose without copying using a row-major submatrix.
func (A SubContiguousRowMajor) T() SubContiguousColMajor {
	return SubContiguousColMajor{A.Cols, A.Rows, A.Stride, A.Elements}
}

func (A SubContiguousRowMajor) Submatrix(r Rect) SubContiguous {
	// Extract from first to last elements.
	i0, j0 := r.Min.I, r.Min.J
	i1, j1 := r.Max.I, r.Max.J
	a := j0*A.Stride + +i0
	b := (j1-1)*A.Stride + (i1 - 1) + 1
	return SubContiguousRowMajor{r.Rows, r.Cols, A.Stride, A.Elements[a:b]}
}
