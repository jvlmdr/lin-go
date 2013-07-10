package mat

import "github.com/jackvalmadre/go-vec"

// Describes a dense matrix with contiguous storage in row-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and row-major enables row slicing and appending.
type ContiguousRowMajor struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elements[i*Cols+j].
	Elements []float64
}

func MakeContiguousRowMajor(rows, cols int) ContiguousRowMajor {
	return ContiguousRowMajor{rows, cols, make([]float64, rows*cols)}
}

func (A ContiguousRowMajor) Size() Size {
	return Size{A.Rows, A.Cols}
}

func (A ContiguousRowMajor) At(i, j int) float64 {
	return A.Elements[i*A.Cols+j]
}

func (A ContiguousRowMajor) Set(i, j int, x float64) {
	A.Elements[i*A.Cols+j] = x
}

func (A ContiguousRowMajor) RowMajor() bool {
	return true
}

func (A ContiguousRowMajor) Array() []float64 {
	return A.Elements
}

// Returns the transpose without copying.
//
// The returned matrix references the same data.
func (A ContiguousRowMajor) T() ContiguousColMajor {
	return ContiguousColMajor{A.Cols, A.Rows, A.Elements}
}

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A ContiguousRowMajor) Resize(s Size) ContiguousRowMajor {
	return A.T().Resize(s.T()).T()
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A ContiguousRowMajor) Slice(i0, i1 int) ContiguousRowMajor {
	return A.T().Slice(i0, i1).T()
}

// Appends a column.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendVector(x vec.Const) ContiguousRowMajor {
	return A.T().AppendVector(x).T()
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendMatrix(B Const) ContiguousRowMajor {
	return A.T().AppendMatrix(T(B)).T()
}

// Appends a column-major matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A ContiguousRowMajor) AppendContiguous(B ContiguousRowMajor) ContiguousRowMajor {
	return A.T().AppendContiguous(B.T()).T()
}

// Selects a submatrix within the contiguous matrix.
func (A ContiguousRowMajor) Submatrix(r Rect) SubContiguousRowMajor {
	return A.T().Submatrix(r.T()).T()
}
