package mat

import "github.com/jackvalmadre/lin-go/vec"

// Submatrix within a contiguous matrix, column-major order.
type ContiguousColMajorSubmat struct {
	Rows int
	Cols int
	Step int
	// The (i, j)-th element is at Elements[j*Step+i].
	Elements []float64
}

func (A ContiguousColMajorSubmat) Size() Size              { return Size{A.Rows, A.Cols} }
func (A ContiguousColMajorSubmat) At(i, j int) float64     { return A.Elements[j*A.Step+i] }
func (A ContiguousColMajorSubmat) Set(i, j int, x float64) { A.Elements[j*A.Step+i] = x }

func (A ContiguousColMajorSubmat) ColMajorArray() []float64 { return A.Elements }
func (A ContiguousColMajorSubmat) Stride() int              { return A.Step }

// Transpose without copying.
func (A ContiguousColMajorSubmat) T() ContiguousRowMajorSubmat {
	return ContiguousRowMajorSubmat(A)
}

func (A ContiguousColMajorSubmat) Submat(r Rect) ContiguousColMajorSubmat {
	// Extract from first to last elements.
	i0, j0 := r.Min.I, r.Min.J
	i1, j1 := r.Max.I, r.Max.J
	a := j0*A.Step + i0
	b := (j1-1)*A.Step + (i1 - 1) + 1
	return ContiguousColMajorSubmat{r.Rows(), r.Cols(), A.Step, A.Elements[a:b]}
}

// Returns MutableVec(A).
func (A ContiguousColMajorSubmat) Vec() vec.Mutable { return MutableVec(A) }

// Returns a mutable column as a slice vector.
func (A ContiguousColMajorSubmat) Col(j int) vec.Slice {
	return ContiguousCol(A, j)
}

// Returns MutableRow(A).
func (A ContiguousColMajorSubmat) Row(i int) vec.Mutable { return MutableRow(A, i) }
