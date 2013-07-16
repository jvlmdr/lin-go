package mat

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/vec"
)

// Evaluates a matrix multiplication element by element.
// Returns a thin wrapper which evaluates the operation on demand.
//
// Caution: It is not a good idea to chain this call!
//	// Very inefficient!
//	Times(A, Times(B, C))
//
//	// Better:
//	Times(A, MakeCopy(Times(B, C)))
// Or even better, check out Multiply().
func Times(A, B Const) Const { return timesExpr{A, B} }

type timesExpr struct{ A, B Const }

func (expr timesExpr) Size() Size { return Size{Rows(expr.A), Cols(expr.B)} }

func (expr timesExpr) At(i, j int) float64 {
	return TimesVec(expr.A, Col(expr.B, j)).At(i)
}

// Multiplies a vector by a matrix.
// Returns a thin wrapper which evaluates the operation on demand.
func TimesVec(A Const, x vec.Const) vec.Const { return timesVecExpr{A, x} }

type timesVecExpr struct {
	A Const
	X vec.Const
}

func (expr timesVecExpr) Size() int { return Rows(expr.A) }

func (expr timesVecExpr) At(i int) float64 {
	n := expr.X.Size()
	var y float64
	for j := 0; j < n; j++ {
		y += expr.A.At(i, j) * expr.X.At(j)
	}
	return y
}

// Returns the horizontal augmentation [A, B].
func Augment(A, B Const) Const {
	if Rows(A) != Rows(B) {
		panic(fmt.Sprintf("Matrices cannot be augmented (%v and %v)", A.Size(), B.Size()))
	}
	return augmentExpr{A, B}
	//return Unvec(vec.Cat(Vec(A), Vec(B)), Rows(A), Cols(A)+Cols(B))
}

type augmentExpr struct{ A, B Const }

func (expr augmentExpr) Size() Size {
	rows, cols := RowsCols(expr.A)
	return Size{rows, cols + Cols(expr.B)}
}

func (expr augmentExpr) At(i, j int) float64 {
	n := Cols(expr.A)
	if j < n {
		return expr.A.At(i, j)
	}
	return expr.B.At(i, j-n)
}

// Returns the vertical stacking [A; B].
func Stack(A, B Const) Const {
	if Cols(A) != Cols(B) {
		panic(fmt.Sprintf("Matrices cannot be stacked (%v and %v)", A.Size(), B.Size()))
	}
	return T(Augment(T(A), T(B)))
}

// Returns the horizontal augmentation [A, B].
func MutableAugment(A, B Mutable) Mutable {
	rows := Rows(A)
	cols := Cols(A) + Cols(B)
	return MutableUnvec(vec.MutableCat(MutableVec(A), MutableVec(B)), rows, cols)
}

// Returns the vertical stacking [A; B].
func MutableStack(A, B Mutable) Mutable {
	return MutableT(MutableAugment(MutableT(A), MutableT(B)))
}
