package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with contiguous storage in row-major order.
type ContigT Contig

// Makes a new rows x cols contiguous matrix.
func MakeContigRowMajor(rows, cols int) ContigT {
	return MakeContig(cols, rows).T()
}

// Copies B into a new contiguous matrix.
func MakeContigRowMajorCopy(B Const) ContigT {
	return MakeContigCopy(T(B)).T()
}

func (A ContigT) Size() Size              { return A.T().Size().T() }
func (A ContigT) At(i, j int) float64     { return A.T().At(j, i) }
func (A ContigT) Set(i, j int, x float64) { A.T().Set(j, i, x) }

func (A ContigT) RowMajorArray() []float64 { return A.T().ColMajorArray() }
func (A ContigT) RowStride() int           { return A.T().ColStride() }

// Transpose without copying.
func (A ContigT) T() Contig { return Contig(A) }

// Modifies the number of rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A ContigT) Reshape(s Size) ContigT {
	return A.T().Reshape(s.T()).T()
}

// Appends a matrix vertically. The number of columns must match.
//
// The returned matrix may reference the same data.
func (A ContigT) RowAppend(B Const) ContigT {
	return A.T().ColAppend(T(B)).T()
}

// Returns a slice of the rows.
//
// The returned matrix references the same data.
func (A ContigT) RowSlice(i0, i1 int) ContigT {
	return A.T().ColSlice(i0, i1).T()
}

// Returns a mutable row as a slice vector.
func (A ContigT) Row(i int) vec.Slice { return A.T().Col(i) }
