package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type ContiguousColMajor struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elements[j*Rows+i].
	Elements []float64
}

// Makes a new rows x cols contiguous matrix.
func MakeContiguousColMajor(rows, cols int) ContiguousColMajor {
	return ContiguousColMajor{rows, cols, make([]float64, rows*cols)}
}

// Copies B into a new contiguous matrix.
func MakeContiguousColMajorCopy(B Const) ContiguousColMajor {
	rows, cols := RowsCols(B)
	A := MakeContiguous(rows, cols)
	Copy(A, B)
	return A
}

func (A ContiguousColMajor) Size() Size              { return Size{A.Rows, A.Cols} }
func (A ContiguousColMajor) At(i, j int) float64     { return A.Elements[j*A.Rows+i] }
func (A ContiguousColMajor) Set(i, j int, x float64) { A.Elements[j*A.Rows+i] = x }

func (A ContiguousColMajor) ColMajorArray() []float64 { return A.Elements }
func (A ContiguousColMajor) Stride() int              { return A.Rows }

// Transpose without copying.
func (A ContiguousColMajor) T() ContiguousRowMajor { return ContiguousRowMajor(A) }

// Returns a vectorization which accesses the array directly.
func (A ContiguousColMajor) Vec() vec.Mutable { return vec.Slice(A.Elements) }

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A ContiguousColMajor) Reshape(s Size) ContiguousColMajor {
	if s.Area() != A.Size().Area() {
		panic("Number of elements must match to resize")
	}
	return ContiguousColMajor{s.Rows, s.Cols, A.Elements}
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A ContiguousColMajor) Slice(j0, j1 int) ContiguousColMajor {
	return ContiguousColMajor{A.Rows, j1 - j0, A.Elements[j0*A.Rows : j1*A.Rows]}
}

// Appends a column.
//
// The returned matrix may reference the same data.
func (A ContiguousColMajor) AppendVector(x vec.Const) ContiguousColMajor {
	if A.Rows != x.Size() {
		panic("Dimension of vector does not match matrix")
	}
	elements := vec.Append(A.Elements, x)
	return ContiguousColMajor{A.Rows, A.Cols + 1, elements}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A ContiguousColMajor) AppendMatrix(B Const) ContiguousColMajor {
	if A.Rows != Rows(B) {
		panic("Dimension of matrices does not match")
	}
	elements := vec.Append(A.Elements, Vec(B))
	return ContiguousColMajor{A.Rows, A.Cols + Cols(B), elements}
}

// Appends a column-major matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A ContiguousColMajor) AppendContiguous(B ContiguousColMajor) ContiguousColMajor {
	if A.Rows != B.Rows {
		panic("Dimension of matrices does not match")
	}
	elements := append(A.Elements, B.Elements...)
	return ContiguousColMajor{A.Rows, A.Cols + B.Cols, elements}
}

// Selects a submatrix within the contiguous matrix.
func (A ContiguousColMajor) Submat(r Rect) ContiguousColMajorSubmat {
	// Extract from first to last element.
	a := r.Min.I + r.Min.J*A.Rows
	b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
	return ContiguousColMajorSubmat{r.Rows(), r.Cols(), A.Rows, A.Elements[a:b]}
}

// Returns MutableColumn(A, j).
func (A ContiguousColMajor) Col(j int) vec.Mutable { return MutableCol(A, j) }

// Returns MutableRow(A, i).
func (A ContiguousColMajor) Row(i int) vec.Mutable { return MutableRow(A, i) }
