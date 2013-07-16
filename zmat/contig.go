package zmat

import "github.com/jackvalmadre/lin-go/zvec"

// Describes a dense matrix with semi-contiguous column-major storage.
// This means that each column is a contiguous array and the columns are regularly spaced, but not necessarily adjacent.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type SemiContiguousColMajor interface {
	Mutable
	ColMajorArray() []complex128
	Stride() int
}

// Makes a new rows x cols contiguous matrix.
func MakeContiguous(rows, cols int) ContiguousColMajor {
	return MakeContiguousColMajor(rows, cols)
}

// Copies B into a new contiguous matrix.
func MakeContiguousCopy(B Const) ContiguousColMajor {
	return MakeContiguousColMajorCopy(B)
}

// Returns a mutable column as a slice vector.
func ContiguousCol(A SemiContiguousColMajor, j int) zvec.Slice {
	a := j * A.Stride()
	b := j*A.Stride() + Rows(A)
	return zvec.Slice(A.ColMajorArray()[a:b])
}

// Selects a submatrix within the contiguous matrix.
func SemiContiguousSubmat(A SemiContiguousColMajor, r Rect) ContiguousColMajorSubmat {
	// Extract from first to last element.
	a := r.Min.J*A.Stride() + r.Min.I
	b := (r.Max.J-1)*A.Stride() + r.Max.I
	elements := A.ColMajorArray()
	return ContiguousColMajorSubmat{r.Rows(), r.Cols(), A.Stride(), elements[a:b]}
}
