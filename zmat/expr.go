package zmat

import "github.com/jackvalmadre/go-vec/zvec"

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

func (expr timesExpr) At(i, j int) complex128 {
	return TimesVec(expr.A, Col(expr.B, j)).At(i)
}

// Multiplies a vector by a matrix.
// Returns a thin wrapper which evaluates the operation on demand.
func TimesVec(A Const, x zvec.Const) zvec.Const { return timesVecExpr{A, x} }

type timesVecExpr struct {
	A Const
	X zvec.Const
}

func (expr timesVecExpr) Size() int { return Rows(expr.A) }

func (expr timesVecExpr) At(i int) complex128 {
	n := expr.X.Size()
	var y complex128
	for j := 0; j < n; j++ {
		y += expr.A.At(i, j) * expr.X.At(j)
	}
	return y
}

// Returns the horizontal augmentation [A, B].
func Augment(A, B Const) Const {
	return Unvec(zvec.Cat(Vec(A), Vec(B)), Rows(A), Cols(A)+Cols(B))
}

// Returns the vertical stacking [A; B].
func Stack(A, B Const) Const {
	return T(Augment(T(A), T(B)))
}

// Returns the horizontal augmentation [A, B].
func MutableAugment(A, B Mutable) Mutable {
	rows := Rows(A)
	cols := Cols(A) + Cols(B)
	return MutableUnvec(zvec.MutableCat(MutableVec(A), MutableVec(B)), rows, cols)
}

// Returns the vertical stacking [A; B].
func MutableStack(A, B Mutable) Mutable {
	return MutableT(MutableAugment(MutableT(A), MutableT(B)))
}
