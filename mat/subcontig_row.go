package mat

import "github.com/jackvalmadre/lin-go/vec"

// Submatrix within a contiguous matrix, row-major order.
type ContiguousRowMajorSubmat ContiguousColMajorSubmat

func (A ContiguousRowMajorSubmat) Size() Size              { return A.T().Size().T() }
func (A ContiguousRowMajorSubmat) At(i, j int) float64     { return A.T().At(j, i) }
func (A ContiguousRowMajorSubmat) Set(i, j int, x float64) { A.T().Set(j, i, x) }

func (A ContiguousRowMajorSubmat) RowMajorArray() []float64 { return A.T().ColMajorArray() }
func (A ContiguousRowMajorSubmat) Stride() int              { return A.T().Stride() }

// Transpose without copying.
func (A ContiguousRowMajorSubmat) T() ContiguousColMajorSubmat {
	return ContiguousColMajorSubmat(A)
}

func (A ContiguousRowMajorSubmat) Submat(r Rect) ContiguousRowMajorSubmat {
	return A.T().Submat(r.T()).T()
}

// Returns MutableVec(A).
func (A ContiguousRowMajorSubmat) Vec() vec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A).
func (A ContiguousRowMajorSubmat) Col(j int) vec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A).
func (A ContiguousRowMajorSubmat) Row(i int) vec.Mutable { return MutableRow(A, i) }
