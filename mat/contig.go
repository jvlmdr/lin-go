package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with semi-contiguous column-major storage.
// This means that each column is a contiguous array and the columns are regularly spaced, but not necessarily adjacent.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type ColMajor interface {
	Mutable
	ColMajorArray() []float64
	Stride() int
}

// Returns a mutable column as a slice vector.
func ContiguousCol(A ColMajor, j int) vec.Slice {
	a := j * A.Stride()
	b := j*A.Stride() + Rows(A)
	return vec.Slice(A.ColMajorArray()[a:b])
}

// Selects a submatrix within the contiguous matrix.
func SemiContiguousSubmat(A ColMajor, r Rect) ContiguousSubmat {
	// Extract from first to last element.
	a := r.Min.J*A.Stride() + r.Min.I
	b := (r.Max.J-1)*A.Stride() + r.Max.I
	elements := A.ColMajorArray()
	return ContiguousSubmat{r.Rows(), r.Cols(), A.Stride(), elements[a:b]}
}
