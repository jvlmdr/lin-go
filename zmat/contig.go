package zmat

// Describes a dense matrix with contiguous storage.
type Contiguous interface {
	Mutable
	RowMajor() bool
	Array() []complex128
}

// Describes a dense matrix with sub-contiguous storage.
//
// Note that Contiguous matrices are also SubContiguous matrices.
type SubContiguous interface {
	Contiguous
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
