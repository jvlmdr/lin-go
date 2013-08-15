package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type Contig struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elems[i + j*Rows].
	// The length of Elems is Rows * Cols.
	Elems []complex128
}

// Makes a new rows x cols contiguous matrix.
func MakeContig(rows, cols int) Contig {
	return Contig{rows, cols, make([]complex128, rows*cols)}
}

// Makes a new rows x cols contiguous matrix.
// The column capacity of the matrix can be specified larger than its size.
func MakeContigCap(rows, cols, colcap int) Contig {
	colcap = max(cols, colcap)
	return Contig{rows, cols, make([]complex128, rows*cols, rows*colcap)}
}

// Copies B into a new contiguous matrix.
func MakeContigCopy(B Const) Contig {
	rows, cols := RowsCols(B)
	A := MakeContig(rows, cols)
	Copy(A, B)
	return A
}

// Creates a column matrix from x.
func ContigMat(x []complex128) Contig {
	return Contig{len(x), 1, x}
}

func (A Contig) Size() Size                 { return Size{A.Rows, A.Cols} }
func (A Contig) At(i, j int) complex128     { return A.Elems[j*A.Rows+i] }
func (A Contig) Set(i, j int, x complex128) { A.Elems[j*A.Rows+i] = x }

// The column capacity using the current number of rows.
func (A Contig) ColCap() int {
	return cap(A.Elems) / A.Rows
}

// Transpose without copying.
func (A Contig) T() ContigT { return ContigT(A) }

func (A Contig) ConstT() Const     { return A.T() }
func (A Contig) MutableT() Mutable { return A.T() }

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

func (A Contig) ConstReshape(s Size) Const     { return A.Reshape(s) }
func (A Contig) MutableReshape(s Size) Mutable { return A.Reshape(s) }

// Turns a contiguous matrix into a stride matrix.
func (A Contig) Stride() Stride {
	return Stride{A.Rows, A.Cols, A.Rows, A.Elems}
}

// Slices the columns.
//
// The returned matrix references the same data.
func (A Contig) ColSlice(j0, j1 int) Contig {
	return Contig{A.Rows, j1 - j0, A.Elems[j0*A.Rows : j1*A.Rows]}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contig) ColAppend(B Const) Contig {
	if A.Rows != Rows(B) {
		panic("different number of rows")
	}

	elems := zvec.Append(A.Elems, Vec(B))
	return Contig{A.Rows, A.Cols + Cols(B), elems}
}

// Returns a vectorization which accesses the array directly.
func (A Contig) Vec() zvec.Slice { return zvec.Slice(A.Elems) }

func (A Contig) ConstVec() zvec.Const     { return A.Vec() }
func (A Contig) MutableVec() zvec.Mutable { return A.Vec() }

// Returns a mutable column as a slice vector.
func (A Contig) Col(j int) zvec.Slice {
	return A.Elems[j*A.Rows : (j+1)*A.Rows]
}

func (A Contig) ConstCol(j int) zvec.Const     { return A.Col(j) }
func (A Contig) MutableCol(j int) zvec.Mutable { return A.Col(j) }

// Slices a rectangle of the matrix, which will not necessarily be contiguous.
func (A Contig) Slice(r Rect) Stride { return A.Stride().Slice(r) }

// Slices a rectangle within the matrix, which will not necessarily be contiguous.
func (A Contig) Submat(r Rect) Stride { return A.Stride().Submat(r) }

func (A Contig) ConstSubmat(r Rect) Const     { return A.Submat(r) }
func (A Contig) MutableSubmat(r Rect) Mutable { return A.Submat(r) }
