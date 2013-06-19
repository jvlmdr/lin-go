package vec

import (
	"github.com/skelterjohn/go.matrix"
)

// Most types should be responsible for providing XxxAsVector themselves.

// Describes a constant vectorized matrix.
type ConstMatrixAsVector struct {
	Matrix   matrix.MatrixRO
	RowMajor bool
}

func (x ConstMatrixAsVector) Size() int {
	return x.Matrix.NumElements()
}

func (x ConstMatrixAsVector) At(i int) float64 {
	r, c := x.Matrix.GetSize()
	u, v := unvectorizeIndex(i, r, c, x.RowMajor)
	return x.Matrix.Get(u, v)
}

// Describes a mutable vectorized matrix.
type MatrixAsVector struct {
	Matrix   matrix.Matrix
	RowMajor bool
}

func (x MatrixAsVector) Size() int {
	return x.Matrix.NumElements()
}

func (x MatrixAsVector) At(i int) float64 {
	r, c := x.Matrix.GetSize()
	u, v := unvectorizeIndex(i, r, c, x.RowMajor)
	return x.Matrix.Get(u, v)
}

func (x MatrixAsVector) Set(i int, a float64) {
	r, c := x.Matrix.GetSize()
	u, v := unvectorizeIndex(i, r, c, x.RowMajor)
	x.Matrix.Set(u, v, a)
}

func unvectorizeIndex(k int, rows, cols int, rowMajor bool) (i, j int) {
	if rowMajor {
		i = k / cols
		j = k % cols
	} else {
		i = k % rows
		j = k / rows
	}
	return
}

func vectorizeIndex(i, j int, rows, cols int, rowMajor bool) int {
	if rowMajor {
		return i*cols + j
	} else {
		return i + j*rows
	}
}
