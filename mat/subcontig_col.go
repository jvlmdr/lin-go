package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a matrix in contiguous memory with a stride between columns.
type Stride struct {
	// LenI
	Rows int
	// LenJ
	Cols int
	// Distance between first elements in adjacent columns.
	Stride int
	// The (i, j)-th element is at Elems[j*Stride+i].
	Elems []float64
}

func MakeStride(rows, cols int) Stride {
	elems := make([]float64, rows*cols)
	return Stride{rows, cols, rows, elems}
}

func MakeStrideCap(rows, cols, rowcap, colcap int) Stride {
	elems := make([]float64, rowcap*cols, rowcap*colcap)
	return Stride{rows, cols, rowcap, elems}
}

func (A Stride) Size() Size              { return Size{A.Rows, A.Cols} }
func (A Stride) At(i, j int) float64     { return A.Elems[j*A.Stride+i] }
func (A Stride) Set(i, j int, x float64) { A.Elems[j*A.Stride+i] = x }

func (A Stride) ColMajorArray() []float64 { return A.Elems }
func (A Stride) ColStride() int           { return A.Stride }

// Transpose without copying.
func (A Stride) T() StrideT { return StrideT(A) }

// Slices the matrix.
// May go beyond the bounds of the matrix but cannot exceed its capacity.
func (A Stride) Slice(r Rect) Stride {
	// Extract from first to last elements.
	a := r.Min.J*A.Stride + r.Min.I
	b := (r.Max.J-1)*A.Stride + (r.Max.I - 1) + 1
	return Stride{r.Rows(), r.Cols(), A.Stride, A.Elems[a:b]}
}

// Slices the bottom right corner of a matrix.
func (A Stride) SliceFrom(i, j int) Stride {
	a := j*A.Stride + i
	return Stride{A.Rows - i, A.Cols - j, A.Stride, A.Elems[a:]}
}

// Slices the top left corner of a matrix.
func (A Stride) SliceTo(i, j int) Stride {
	// Extract from first to last elements.
	b := (j-1)*A.Stride + (i - 1) + 1
	return Stride{i, j, A.Stride, A.Elems[:b]}
}

// Can the matrix be sliced without re-allocating?
//
// Provides the equivalent functionality to n <= cap(x),
// though the relationship is more complex for matrices with a stride.
func (A Stride) InCap(r Rect) bool {
	size := r.Size()
	if size.Rows > A.Stride {
		return false
	}
	// Check location of last element.
	b := (size.Cols-1)*A.Stride + (size.Rows - 1) + 1
	return b <= cap(A.Elems)
}

// Can the matrix be grown to include point (i, j)?
func (A Stride) InCapTo(i, j int) bool {
	if i > A.Stride {
		// More rows than stride.
		return false
	}
	// Check location of last element.
	b := (j-1)*A.Stride + (i - 1) + 1
	return b <= cap(A.Elems)
}

// The column capacity using the current number of rows.
func (A Stride) ColCap() int {
	return cap(A.Elems) / A.Rows
}

// The row capacity using the current number of columns.
func (A Stride) RowCap() int {
	// Start of last column.
	b := (A.Cols - 1) * A.Stride
	return min(A.Stride, cap(A.Elems)-b)
}

// Minimum column capacity occurs when number of rows equals the stride.
func (A Stride) minColCap() int {
	return cap(A.Elems) / A.Stride
}

// Maximum column capacity occurs when number of rows is one.
func (A Stride) maxColCap() int {
	return (cap(A.Elems)-1)/A.Stride + 1
}

func (A Stride) AppendCols(B Const) Stride {
	if A.Rows != Rows(B) {
		panic("different number of rows")
	}

	cols := A.Cols + Cols(B)
	// Re-allocate if necessary.
	if !A.InCapTo(A.Rows, cols) {
		// Keep same stride. At least double column capacity.
		rowcap := A.Stride
		colcap := max(2*A.maxColCap(), cols)
		X := MakeStrideCap(A.Rows, A.Cols, rowcap, colcap)
		// Copy to matrix with same size, greater capacity.
		Copy(X, A)
		A = X
	}
	// Grow size of matrix.
	X := A.SliceTo(A.Rows, cols)

	// Copy new part.
	Copy(X.SliceFrom(0, A.Cols), B)
	return X
}

func (A Stride) AppendRows(B Const) Stride {
	if A.Cols != Cols(B) {
		panic("different number of cols")
	}

	rows := A.Rows + Rows(B)
	// Re-allocate if necessary.
	if !A.InCapTo(rows, A.Cols) {
		// Increase stride. Preserve maximum column capacity.
		rowcap := max(2*A.Stride, rows)
		colcap := A.maxColCap()
		X := MakeStrideCap(A.Rows, A.Cols, rowcap, colcap)
		// Copy to matrix with same size, greater capacity.
		Copy(X, A)
		A = X
	}
	// Grow size of matrix.
	X := A.SliceTo(rows, A.Cols)

	// Copy new part.
	Copy(X.SliceFrom(A.Rows, 0), B)
	return X
}

// Returns MutableVec(A).
func (A Stride) Vec() vec.Mutable { return MutableVec(A) }

// Returns a mutable column as a slice vector.
func (A Stride) Col(j int) vec.Slice { return ContigCol(A, j) }

// Returns MutableRow(A).
func (A Stride) Row(i int) vec.Mutable { return MutableRow(A, i) }
