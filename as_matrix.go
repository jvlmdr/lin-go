package vec

import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
)

// Wraps a vec.Const to behave as an nx1 MatrixRO.
type AsMatrix struct {
	Vector Const
}

func (x AsMatrix) Nil() bool {
	return x.Vector == nil
}

func (x AsMatrix) Rows() int {
	return x.Vector.Size()
}

func (x AsMatrix) Cols() int {
	return 1
}

func (x AsMatrix) NumElements() int {
	return x.Vector.Size()
}

func (x AsMatrix) GetSize() (int, int) {
	return x.Vector.Size(), 1
}

func (x AsMatrix) Get(i, j int) float64 {
	n := x.Vector.Size()
	if j != 0 || !(i >= 0 && i < n) {
		problem := fmt.Sprintf("Out of range: Tried to access (%d, %d) in %dx1 matrix", i, j, n)
		panic(problem)
	}
	return x.Vector.At(i)
}

func (x AsMatrix) Plus(y matrix.MatrixRO) (matrix.Matrix, error) {
	return x.DenseMatrix().Plus(y)
}

func (x AsMatrix) Minus(y matrix.MatrixRO) (matrix.Matrix, error) {
	return x.DenseMatrix().Minus(y)
}

func (x AsMatrix) Times(y matrix.MatrixRO) (matrix.Matrix, error) {
	return x.DenseMatrix().Times(y)
}

func (x AsMatrix) Det() float64 {
	panic("Not a square matrix")
}

func (x AsMatrix) Trace() float64 {
	panic("Not a square matrix")
}

func (x AsMatrix) String() string {
	return String(x.Vector)
}

func (x AsMatrix) DenseMatrix() *matrix.DenseMatrix {
	n := x.Vector.Size()
	y := matrix.Zeros(n, 1)
	// RowMajor is unimportant for a column matrix.
	CopyTo(MatrixAsVector{y, false}, x.Vector)
	return y
}

func (x AsMatrix) SparseMatrix() *matrix.SparseMatrix {
	panic("Unimplemented")
}
