package mat

// Describes a dense matrix with semi-contiguous column-major storage.
// This means that each column is a contiguous array and the columns are regularly spaced, but not necessarily adjacent.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type ColMajor interface {
	Mutable
	ColMajorArray() []float64
	ColStride() int
}
