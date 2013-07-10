package mat

import "github.com/jackvalmadre/go-vec"

// Submatrix within a contiguous matrix, column-major order.
type SubContiguousColMajor struct {
	Rows   int
	Cols   int
	Stride int
	// The (i, j)-th element is at Elements[j*Stride+i].
	Elements []float64
}

func (A SubContiguousColMajor) Size() Size {
	return Size{A.Rows, A.Cols}
}

func (A SubContiguousColMajor) At(i, j int) float64 {
	return A.Elements[j*A.Stride+i]
}

func (A SubContiguousColMajor) Set(i, j int, x float64) {
	A.Elements[j*A.Stride+i] = x
}

func (A SubContiguousColMajor) RowMajor() bool {
	return false
}

func (A SubContiguousColMajor) Array() []float64 {
	return A.Elements
}

// Returns the transpose without copying using a row-major submatrix.
func (A SubContiguousColMajor) T() SubContiguousRowMajor {
	return SubContiguousRowMajor{A.Cols, A.Rows, A.Stride, A.Elements}
}

func (A SubContiguousColMajor) Submatrix(r Rect) SubContiguous {
	// Extract from first to last elements.
	i0, j0 := r.Min.I, r.Min.J
	i1, j1 := r.Max.I, r.Max.J
	a := j0*A.Stride + i0
	b := (j1-1)*A.Stride + (i1 - 1) + 1
	return SubContiguousColMajor{r.Rows, r.Cols, A.Stride, A.Elements[a:b]}
}
