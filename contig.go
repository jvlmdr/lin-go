package mat

import "github.com/jackvalmadre/go-vec"

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type Contiguous struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elements[j*Rows+i].
	Elements []float64
}

// Makes a new rows x cols contiguous matrix.
func MakeContiguous(rows, cols int) Contiguous {
	return Contiguous{rows, cols, make([]float64, rows*cols)}
}

// Copies B into a new contiguous matrix.
func MakeContiguousCopy(B Const) Contiguous {
	rows, cols := RowsCols(B)
	A := MakeContiguous(rows, cols)
	Copy(A, B)
	return A
}

func (A Contiguous) Size() Size              { return Size{A.Rows, A.Cols} }
func (A Contiguous) At(i, j int) float64     { return A.Elements[j*A.Rows+i] }
func (A Contiguous) Set(i, j int, x float64) { A.Elements[j*A.Rows+i] = x }

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A Contiguous) Resize(s Size) Contiguous {
	if s.Area() != A.Size().Area() {
		panic("Number of elements must match to resize")
	}
	return Contiguous{s.Rows, s.Cols, A.Elements}
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A Contiguous) Slice(j0, j1 int) Contiguous {
	return Contiguous{A.Rows, j1 - j0, A.Elements[j0*A.Rows : j1*A.Rows]}
}

// Appends a column.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendVector(x vec.Const) Contiguous {
	if A.Rows != x.Size() {
		panic("Dimension of vector does not match matrix")
	}
	elements := vec.AppendToSlice(A.Elements, x)
	return Contiguous{A.Rows, A.Cols + 1, elements}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendMatrix(B Const) Contiguous {
	if A.Rows != Rows(B) {
		panic("Dimension of matrices does not match")
	}
	elements := vec.AppendToSlice(A.Elements, Vec(B))
	return Contiguous{A.Rows, A.Cols + Cols(B), elements}
}

// Appends a column-major matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendContiguous(B Contiguous) Contiguous {
	if A.Rows != B.Rows {
		panic("Dimension of matrices does not match")
	}
	elements := append(A.Elements, B.Elements...)
	return Contiguous{A.Rows, A.Cols + B.Cols, elements}
}

// Selects a submatrix within the contiguous matrix.
func (A Contiguous) Submatrix(r Rect) SubContiguous {
	// Extract from first to last element.
	a := r.Min.I + r.Min.J*A.Rows
	b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
	return SubContiguous{r.Rows(), r.Cols(), A.Rows, A.Elements[a:b]}
}

// Returns a vectorization which accesses the array directly.
func (A Contiguous) Vec() vec.Mutable { return contiguousAsVector(A) }

// Specialized vectorization which accesses the array directly.
type contiguousAsVector Contiguous

func (x contiguousAsVector) Size() int            { return x.Rows * x.Cols }
func (x contiguousAsVector) At(i int) float64     { return x.Elements[i] }
func (x contiguousAsVector) Set(i int, v float64) { x.Elements[i] = v }

// Returns MutableT(A).
func (A Contiguous) T() Mutable { return MutableT(A) }
