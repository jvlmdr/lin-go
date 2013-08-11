package mat

import "github.com/jackvalmadre/lin-go/vec"

// Describes a dense matrix with semi-contiguous column-major storage.
// This means that each column is a contiguous array and the columns are regularly spaced, but not necessarily adjacent.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type ColMajor interface {
	Mutable
	ColMajorArray() []float64
	ColStride() int
}

// Returns a mutable column as a slice vector.
func ContigCol(A ColMajor, j int) vec.Slice {
	a := j * A.ColStride()
	b := j*A.ColStride() + Rows(A)
	return vec.Slice(A.ColMajorArray()[a:b])
}

// Selects a submatrix within the contiguous matrix.
func ColMajorSubmat(A ColMajor, r Rect) Stride {
	// Extract from first to last element.
	a := r.Min.J*A.ColStride() + r.Min.I
	b := (r.Max.J-1)*A.ColStride() + r.Max.I
	elements := A.ColMajorArray()
	return Stride{r.Rows(), r.Cols(), A.ColStride(), elements[a:b]}
}
