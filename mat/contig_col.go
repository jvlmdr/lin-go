package mat

import (
	"github.com/jackvalmadre/lin-go/vec"

	"bytes"
	"fmt"
	"io"
)

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

// Creates a column matrix from x.
func FromSlice(x []float64) Contiguous {
	return Contiguous{len(x), 1, x}
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

func (A Contiguous) ColMajorArray() []float64 { return A.Elements }
func (A Contiguous) Stride() int              { return A.Rows }

func (A Contiguous) String() string {
	var b bytes.Buffer
	if err := A.Fprintf(&b, "%.4g"); err != nil {
		panic(err)
	}
	return b.String()
}

func (A Contiguous) Fprintf(w io.Writer, format string) error {
	// One pass to get max length.
	var maxlen int
	for j := 0; j < A.Cols; j++ {
		for i := 0; i < A.Rows; i++ {
			s := fmt.Sprintf(format, A.At(i, j))
			if len(s) > maxlen {
				maxlen = len(s)
			}
		}
	}

	var b bytes.Buffer
	b.WriteString("[")

	for j := 0; j < A.Cols; j++ {
		for i := 0; i < A.Rows; i++ {
			// Convert to string, pad with spaces, write.
			s := fmt.Sprintf(format, A.At(i, j))
			for n := len(s); n < maxlen; n++ {
				b.WriteString(" ")
			}
			b.WriteString(s)

			if i < A.Rows - 1 {
				b.WriteString(", ")
			} else if j < A.Cols - 1 {
				b.WriteString(",\n ")
			} else {
				b.WriteString("]")
			}

			// Copy to output.
			_, err := io.Copy(w, &b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Transpose without copying.
func (A Contiguous) T() ContiguousRowMajor { return ContiguousRowMajor(A) }

// Returns a vectorization which accesses the array directly.
func (A Contiguous) Vec() vec.Mutable { return vec.Slice(A.Elements) }

// Modifies the rows and columns of a contiguous matrix.
// The number of elements must remain constant.
//
// The returned matrix references the same data.
func (A Contiguous) Reshape(s Size) Contiguous {
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
	if A.Rows != x.Len() {
		panic("Dimension of vector does not match matrix")
	}
	elements := vec.Append(A.Elements, x)
	return Contiguous{A.Rows, A.Cols + 1, elements}
}

// Appends a matrix horizontally. The number of rows must match.
//
// The returned matrix may reference the same data.
func (A Contiguous) AppendMatrix(B Const) Contiguous {
	if A.Rows != Rows(B) {
		panic("Dimension of matrices does not match")
	}
	elements := vec.Append(A.Elements, Vec(B))
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
func (A Contiguous) Submat(r Rect) ContiguousSubmat {
	// Extract from first to last element.
	a := r.Min.I + r.Min.J*A.Rows
	b := (r.Max.I - 1) + (r.Max.J-1)*A.Rows + 1
	return ContiguousSubmat{r.Rows(), r.Cols(), A.Rows, A.Elements[a:b]}
}

// Returns a mutable column as a slice vector.
func (A Contiguous) Col(j int) vec.Slice {
	return ContiguousCol(A, j)
}

// Returns MutableRow(A, i).
func (A Contiguous) Row(i int) vec.Mutable { return MutableRow(A, i) }
