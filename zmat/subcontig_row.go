package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Submatrix within a contiguous matrix, row-major order.
type SubContiguousRowMajor SubContiguousColMajor

func (A SubContiguousRowMajor) Size() Size                 { return A.T().Size().T() }
func (A SubContiguousRowMajor) At(i, j int) complex128     { return A.T().At(j, i) }
func (A SubContiguousRowMajor) Set(i, j int, x complex128) { A.T().Set(j, i, x) }

func (A SubContiguousRowMajor) RowMajor() bool      { return true }
func (A SubContiguousRowMajor) Array() []complex128 { return A.T().Array() }
func (A SubContiguousRowMajor) Stride() int         { return A.T().Stride() }

// Transpose without copying.
func (A SubContiguousRowMajor) T() SubContiguousColMajor {
	return SubContiguousColMajor(A)
}

func (A SubContiguousRowMajor) Submat(r Rect) SubContiguousRowMajor {
	return A.T().Submat(r.T()).T()
}

// Returns MutableVec(A).
func (A SubContiguousRowMajor) Vec() zvec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A).
func (A SubContiguousRowMajor) Col(j int) zvec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A).
func (A SubContiguousRowMajor) Row(i int) zvec.Mutable { return MutableRow(A, i) }
