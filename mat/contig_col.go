package mat

import (
	"github.com/jackvalmadre/lin-go/vec"

	//"bytes"
	//"fmt"
	//"io"
)

// Describes a dense matrix with contiguous storage in column-major order.
//
// Being contiguous enables reshaping.
// Being contiguous and column-major enables column slicing and appending.
type Contig struct {
	Rows int
	Cols int
	// The (i, j)-th element is at Elems[i + j*Rows].
	Elems []float64
}

// Makes a new rows x cols contiguous matrix.
func MakeContig(rows, cols int) Contig {
	return Contig{rows, cols, make([]float64, rows*cols)}
}

// Creates a column matrix from x.
func FromSlice(x []float64) Contig {
	return Contig{len(x), 1, x}
}

// Copies B into a new contiguous matrix.
func MakeContigCopy(B Const) Contig {
	rows, cols := RowsCols(B)
	A := MakeContig(rows, cols)
	Copy(A, B)
	return A
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

//	func (A Contig) String() string {
//		var b bytes.Buffer
//		if err := A.Fprintf(&b, "%.4g"); err != nil {
//			panic(err)
//		}
//		return b.String()
//	}

//	func (A Contig) Fprintf(w io.Writer, format string) error {
//		// One pass to get max length.
//		var maxlen int
//		for j := 0; j < A.Cols; j++ {
//			for i := 0; i < A.Rows; i++ {
//				s := fmt.Sprintf(format, A.At(i, j))
//				if len(s) > maxlen {
//					maxlen = len(s)
//				}
//			}
//		}
//
//		var b bytes.Buffer
//		b.WriteString("[")
//
//		for j := 0; j < A.Cols; j++ {
//			for i := 0; i < A.Rows; i++ {
//				// Convert to string, pad with spaces, write.
//				s := fmt.Sprintf(format, A.At(i, j))
//				for n := len(s); n < maxlen; n++ {
//					b.WriteString(" ")
//				}
//				b.WriteString(s)
//
//				if i < A.Rows - 1 {
//					b.WriteString(", ")
//				} else if j < A.Cols - 1 {
//					b.WriteString(",\n ")
//				} else {
//					b.WriteString("]")
//				}
//
//				// Copy to output.
//				_, err := io.Copy(w, &b)
//				if err != nil {
//					return err
//				}
//			}
//		}
//		return nil
//	}

// Modifies the rows and columns of a contiguous matrix.
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
func (A Contig) SliceCols(j0, j1 int) Contig {
	return Contig{A.Rows, j1 - j0, A.Elems[j0*A.Rows : j1*A.Rows]}
}

// Selects a submatrix within the contiguous matrix.
func (A Contig) Slice(r Rect) Stride {
	// Extract from first to last element.
	a := r.Min.I + r.Min.J*A.Rows
	b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
	return Stride{r.Rows(), r.Cols(), A.Rows, A.Elems[a:b]}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contig) AppendCols(B Const) Contig {
	if A.Rows != Rows(B) {
		panic("different number of rows")
	}

	elems := vec.Append(A.Elems, Vec(B))
	return Contig{A.Rows, A.Cols + Cols(B), elems}
}

// Returns a vectorization which accesses the array directly.
func (A Contig) Vec() vec.Mutable { return vec.Slice(A.Elems) }

// Returns a mutable column as a slice vector.
func (A Contig) Col(j int) vec.Slice {
	return ContigCol(A, j)
}

// Returns MutableRow(A, i).
func (A Contig) Row(i int) vec.Mutable { return MutableRow(A, i) }
