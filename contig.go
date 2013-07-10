package mat

import "github.com/jackvalmadre/go-vec"

type SubContiguous interface {
	Mutable
	RowMajor() bool
	Array() []float64
	Submatrix(Rect) SubContiguous
	//Slice(vec.Const) SubContiguous
}

// There is no ConstContiguous matrix since it's hard to imagine a case where a
// matrix can fulfill all functionality except Set().
type Contiguous interface {
	SubContiguous
	Reshape(rows, cols int) Contiguous
	//AppendVector(vec.Const) Contiguous
	//AppendMatrix(Const) Contiguous
}

func MakeGeneralContiguous(rows, cols int, rowmajor bool) MutableContiguous {
	if rowmajor {
		return MakeContiguousRowMajor(rows, cols)
	}
	return MakeContiguousColMajor(rows, cols)
}
