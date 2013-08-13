package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Describes a matrix in contiguous memory with a stride between columns.
type Stride struct {
	// LenI
	Rows int
	// LenJ
	Cols int
	// Distance between first elements in adjacent columns.
	Stride int
	// The (i, j)-th element is at Elems[j*Stride+i].
	// The length of Elems is between Rows * Cols and Stride * Cols inclusive.
	Elems []complex128
}

func MakeStride(rows, cols int) Stride {
	elems := make([]complex128, rows*cols)
	return Stride{rows, cols, rows, elems}
}

func MakeStrideCap(rows, cols, rowcap, colcap int) Stride {
	elems := make([]complex128, rowcap*cols, rowcap*colcap)
	return Stride{rows, cols, rowcap, elems}
}

func (A Stride) Size() Size                 { return Size{A.Rows, A.Cols} }
func (A Stride) At(i, j int) complex128     { return A.Elems[j*A.Stride+i] }
func (A Stride) Set(i, j int, x complex128) { A.Elems[j*A.Stride+i] = x }

func (A Stride) ColMajorArray() []complex128 { return A.Elems }
func (A Stride) ColStride() int              { return A.Stride }

// Transpose without copying.
func (A Stride) T() StrideT { return StrideT(A) }

func (A Stride) ConstT() Const     { return A.T() }
func (A Stride) MutableT() Mutable { return A.T() }

// Turns a stride matrix into a contiguous matrix.
// Panics if the stride is not equal to the number of rows.
func (A Stride) Contig() Contig {
	if A.Rows != A.Stride {
		panic("number of rows does not match stride")
	}
	return Contig{A.Rows, A.Cols, A.Elems}
}

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

// Slices a submatrix within the bounds of the matrix.
func (A Stride) Submat(r Rect) Stride {
	if r.Max.I >= A.Rows || r.Max.J >= A.Cols {
		panic("rectangle is outside range of matrix")
	}
	return A.Slice(r)
}

func (A Stride) ConstSubmat(r Rect) Const     { return A.Submat(r) }
func (A Stride) MutableSubmat(r Rect) Mutable { return A.Submat(r) }

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
	// Find start of last column.
	b := (A.Cols - 1) * A.Stride
	// Limited by stride.
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

func (A Stride) ColAppend(B Const) Stride {
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

func (A Stride) RowAppend(B Const) Stride {
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

// Returns a mutable column as a slice vector.
func (A Stride) Col(j int) zvec.Slice {
	return A.Elems[j*A.Stride : j*A.Stride+A.Rows]
}

func (A Stride) ConstCol(j int) zvec.Const     { return A.Col(j) }
func (A Stride) MutableCol(j int) zvec.Mutable { return A.Col(j) }
