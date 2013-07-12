package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with contiguous storage in row-major order.
type ContiguousRowMajor ContiguousColMajor

// Makes a new rows x cols contiguous matrix.
func MakeContiguousRowMajor(rows, cols int) ContiguousRowMajor {
	return MakeContiguousColMajor(cols, rows).T()
}

// Copies B into a new contiguous matrix.
func MakeContiguousRowMajorCopy(B Const) ContiguousRowMajor {
	return MakeContiguousColMajorCopy(T(B)).T()
}

func (A ContiguousRowMajor) Size() Size              { return A.T().Size().T() }
func (A ContiguousRowMajor) At(i, j int) float64     { return A.T().At(j, i) }
func (A ContiguousRowMajor) Set(i, j int, x float64) { A.T().Set(j, i, x) }

func (A ContiguousRowMajor) RowMajor() bool   { return true }
func (A ContiguousRowMajor) Array() []float64 { return A.T().Array() }
func (A ContiguousRowMajor) Stride() int      { return A.T().Stride() }

// Transpose without copying.
func (A ContiguousRowMajor) T() ContiguousColMajor { return ContiguousColMajor(A) }

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A ContiguousRowMajor) Reshape(s Size) ContiguousRowMajor {
	return A.T().Reshape(s.T()).T()
}

// Appends a row.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendVector(x vec.Const) ContiguousRowMajor {
	return A.T().AppendVector(x).T()
}

// Appends a matrix vertically. The number of columns must match.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendMatrix(B Const) ContiguousRowMajor {
	return A.T().AppendMatrix(T(B)).T()
}

// Appends a row-major matrix vertically. The number of columns must match.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendContiguous(B ContiguousRowMajor) ContiguousRowMajor {
	return A.T().AppendContiguous(B.T()).T()
}

// Returns a slice of the rows.
//
// The returned matrix references the same data.
func (A ContiguousRowMajor) Slice(i0, i1 int) ContiguousRowMajor {
	return A.T().Slice(i0, i1).T()
}

// Selects a submatrix within the contiguous matrix.
func (A ContiguousRowMajor) Submat(r Rect) SubContiguousRowMajor {
	return A.T().Submat(r.T()).T()
}

// Returns MutableVec(A).
func (A ContiguousRowMajor) Vec() vec.Mutable { return MutableVec(A) }

// Returns MutableColumn(A, j).
func (A ContiguousRowMajor) Col(j int) vec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A, i).
func (A ContiguousRowMajor) Row(i int) vec.Mutable { return MutableRow(A, i) }
