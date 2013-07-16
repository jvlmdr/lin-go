package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Submatrix within a contiguous matrix, row-major order.
type ContiguousRowMajorSubmat ContiguousSubmat

func (A ContiguousRowMajorSubmat) Size() Size                 { return A.T().Size().T() }
func (A ContiguousRowMajorSubmat) At(i, j int) complex128     { return A.T().At(j, i) }
func (A ContiguousRowMajorSubmat) Set(i, j int, x complex128) { A.T().Set(j, i, x) }

func (A ContiguousRowMajorSubmat) RowMajorArray() []complex128 { return A.T().ColMajorArray() }
func (A ContiguousRowMajorSubmat) Stride() int                 { return A.T().Stride() }

// Transpose without copying.
func (A ContiguousRowMajorSubmat) T() ContiguousSubmat {
	return ContiguousSubmat(A)
}

func (A ContiguousRowMajorSubmat) Submat(r Rect) ContiguousRowMajorSubmat {
	return A.T().Submat(r.T()).T()
}

// Returns MutableVec(A).
func (A ContiguousRowMajorSubmat) Vec() zvec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A).
func (A ContiguousRowMajorSubmat) Col(j int) zvec.Mutable { return MutableCol(A, j) }

// Returns a mutable row as a slice vector.
func (A ContiguousRowMajorSubmat) Row(i int) zvec.Slice { return A.T().Col(i) }
