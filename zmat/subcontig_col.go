package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Submatrix within a contiguous matrix, column-major order.
type ContiguousColMajorSubmat struct {
	Rows int
	Cols int
	Step int
	// The (i, j)-th element is at Elements[j*Step+i].
	Elements []complex128
}

func (A ContiguousColMajorSubmat) Size() Size                 { return Size{A.Rows, A.Cols} }
func (A ContiguousColMajorSubmat) At(i, j int) complex128     { return A.Elements[j*A.Step+i] }
func (A ContiguousColMajorSubmat) Set(i, j int, x complex128) { A.Elements[j*A.Step+i] = x }

func (A ContiguousColMajorSubmat) ColMajorArray() []complex128 { return A.Elements }
func (A ContiguousColMajorSubmat) Stride() int                 { return A.Step }

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
func (A ContiguousColMajorSubmat) Vec() zvec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A).
func (A ContiguousColMajorSubmat) Col(j int) zvec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A).
func (A ContiguousColMajorSubmat) Row(i int) zvec.Mutable { return MutableRow(A, i) }
