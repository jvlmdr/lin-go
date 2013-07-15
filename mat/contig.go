package mat

// Describes a dense matrix with semi-contiguous column-major storage.
// This means that each column is a contiguous array and the columns are regularly spaced, but not necessarily adjacent.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type SemiContiguousColMajor interface {
	Mutable
	ColMajorArray() []float64
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
