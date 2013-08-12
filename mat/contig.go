package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type Contig struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elems[i + j*Rows].
	// The length of Elems is Rows * Cols.
	Elems []float64
}

// Makes a new rows x cols contiguous matrix.
func MakeContig(rows, cols int) Contig {
	return Contig{rows, cols, make([]float64, rows*cols)}
}

// Makes a new rows x cols contiguous matrix.
// The column capacity of the matrix can be specified larger than its size.
func MakeContigCap(rows, cols, colcap int) Contig {
	colcap = max(cols, colcap)
	return Contig{rows, cols, make([]float64, rows*cols, rows*colcap)}
}

// Copies B into a new contiguous matrix.
func MakeContigCopy(B Const) Contig {
	rows, cols := RowsCols(B)
	A := MakeContig(rows, cols)
	Copy(A, B)
	return A
}

// Creates a column matrix from x.
func ContigMat(x []float64) Contig {
	return Contig{len(x), 1, x}
}

func (A Contig) Size() Size              { return Size{A.Rows, A.Cols} }
func (A Contig) At(i, j int) float64     { return A.Elems[j*A.Rows+i] }
func (A Contig) Set(i, j int, x float64) { A.Elems[j*A.Rows+i] = x }

func (A Contig) ColMajorArray() []float64 { return A.Elems }
func (A Contig) ColStride() int           { return A.Rows }

// Transpose without copying.
func (A Contig) T() ContigT { return ContigT(A) }

// The column capacity using the current number of rows.
func (A Contig) ColCap() int {
	return cap(A.Elems) / A.Rows
}

// Modifies the number of rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A Contig) Reshape(s Size) Contig {
	if s.Area() != A.Size().Area() {
		panic("different number of elements")
	}
	return Contig{s.Rows, s.Cols, A.Elems}
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A Contig) ColSlice(j0, j1 int) Contig {
	return Contig{A.Rows, j1 - j0, A.Elems[j0*A.Rows : j1*A.Rows]}
}

// Turns a contiguous matrix into a stride matrix.
func (A Contig) Stride() Stride {
	return Stride{A.Rows, A.Cols, A.Rows, A.Elems}
}

//	// Selects a submatrix within the contiguous matrix.
//	func (A Contig) Slice(r Rect) Stride {
//		// Extract from first to last element.
//		a := r.Min.I + r.Min.J*A.Rows
//		b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
//		return Stride{r.Rows(), r.Cols(), A.Rows, A.Elems[a:b]}
//	}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contig) ColAppend(B Const) Contig {
	if A.Rows != Rows(B) {
		panic("different number of rows")
	}

	elems := vec.Append(A.Elems, Vec(B))
	return Contig{A.Rows, A.Cols + Cols(B), elems}
}

// Returns a vectorization which accesses the array directly.
func (A Contig) Vec() vec.Slice { return vec.Slice(A.Elems) }

// Returns a mutable column as a slice vector.
func (A Contig) Col(j int) vec.Slice { return ContigCol(A, j) }
