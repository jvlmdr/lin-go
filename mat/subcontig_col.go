package mat

import "github.com/jackvalmadre/lin-go/vec"

// Submatrix within a contiguous matrix, column-major order.
type SubContiguousColMajor struct {
	Rows int
	Cols int
	Step int
	// The (i, j)-th element is at Elements[j*Step+i].
	Elements []float64
}

func (A SubContiguousColMajor) Size() Size              { return Size{A.Rows, A.Cols} }
func (A SubContiguousColMajor) At(i, j int) float64     { return A.Elements[j*A.Step+i] }
func (A SubContiguousColMajor) Set(i, j int, x float64) { A.Elements[j*A.Step+i] = x }

func (A SubContiguousColMajor) RowMajor() bool   { return false }
func (A SubContiguousColMajor) Array() []float64 { return A.Elements }
func (A SubContiguousColMajor) Stride() int      { return A.Step }

// Transpose without copying.
func (A SubContiguousColMajor) T() SubContiguousRowMajor {
	return SubContiguousRowMajor(A)
}

func (A SubContiguousColMajor) Submat(r Rect) SubContiguousColMajor {
	// Extract from first to last elements.
	i0, j0 := r.Min.I, r.Min.J
	i1, j1 := r.Max.I, r.Max.J
	a := j0*A.Step + i0
	b := (j1-1)*A.Step + (i1 - 1) + 1
	return SubContiguousColMajor{r.Rows(), r.Cols(), A.Step, A.Elements[a:b]}
}

// Returns MutableVec(A).
func (A SubContiguousColMajor) Vec() vec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A).
func (A SubContiguousColMajor) Col(j int) vec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A).
func (A SubContiguousColMajor) Row(i int) vec.Mutable { return MutableRow(A, i) }
