package mat

import "github.com/jackvalmadre/lin-go/vec"

// Submatrix within a contiguous matrix, row-major order.
type StrideT Stride

func (A StrideT) Size() Size              { return A.T().Size().T() }
func (A StrideT) At(i, j int) float64     { return A.T().At(j, i) }
func (A StrideT) Set(i, j int, x float64) { A.T().Set(j, i, x) }

func (A StrideT) RowMajorArray() []float64 { return A.T().ColMajorArray() }
func (A StrideT) RowStride() int           { return A.T().ColStride() }

// Transpose without copying.
func (A StrideT) T() Stride { return Stride(A) }

func (A StrideT) ConstT() Const     { return A.T() }
func (A StrideT) MutableT() Mutable { return A.T() }

// See Stride.Slice().
func (A StrideT) Slice(r Rect) StrideT {
	return A.T().Slice(r.T()).T()
}

// See Stride.Submat().
func (A StrideT) Submat(r Rect) StrideT {
	return A.T().Submat(r.T()).T()
}

func (A StrideT) ConstSubmat(r Rect) Const     { return A.Submat(r) }
func (A StrideT) MutableSubmat(r Rect) Mutable { return A.Submat(r) }

// See Stride.SliceFrom().
func (A StrideT) SliceFrom(i, j int) StrideT {
	return A.T().SliceFrom(j, i).T()
}

// See Stride.SliceTo().
func (A StrideT) SliceTo(i, j int) StrideT {
	return A.T().SliceTo(j, i).T()
}

// See Stride.InCap().
func (A StrideT) InCap(r Rect) bool {
	return A.T().InCap(r.T())
}

// See Stride.InCapTo().
func (A StrideT) InCapTo(i, j int) bool {
	return A.T().InCapTo(j, i)
}

// See Stride.RowAppend().
func (A StrideT) RowAppend(B Const) StrideT {
	return A.T().ColAppend(T(B)).T()
}

// See Stride.ColAppend().
func (A StrideT) ColAppend(B Const) StrideT {
	return A.T().RowAppend(T(B)).T()
}

// Returns a mutable row as a slice vector.
func (A StrideT) Row(i int) vec.Slice { return A.T().Col(i) }

func (A StrideT) ConstRow(i int) vec.Const     { return A.Row(i) }
func (A StrideT) MutableRow(i int) vec.Mutable { return A.Row(i) }
