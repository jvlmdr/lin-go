package mat

// This file contains operations which modify the shape of a matrix.

type ConstTer interface {
	ConstT() Const
}

type MutableTer interface {
	MutableT() Mutable
}

// Returns a thin wrapper which transposes a constant matrix.
// If A is a ConstTer, returns A.ConstT().
func T(A Const) Const {
	if A, ok := A.(ConstTer); ok {
		return A.ConstT()
	}
	return tExpr{A}
}

type tExpr struct{ Matrix Const }

func (expr tExpr) Size() Size          { return expr.Matrix.Size().T() }
func (expr tExpr) At(i, j int) float64 { return expr.Matrix.At(j, i) }

// Returns a thin wrapper which transposes a mutable matrix.
// If A is a MutableTer, returns A.MutableT().
func MutableT(A Mutable) Mutable {
	if A, ok := A.(MutableTer); ok {
		return A.MutableT()
	}
	return mutableTExpr{A}
}

type mutableTExpr struct{ Matrix Mutable }

func (expr mutableTExpr) Size() Size              { return expr.Matrix.Size().T() }
func (expr mutableTExpr) At(i, j int) float64     { return expr.Matrix.At(j, i) }
func (expr mutableTExpr) Set(i, j int, v float64) { expr.Matrix.Set(j, i, v) }

type ConstSubmatter interface {
	ConstSubmat(Rect) Const
}

type MutableSubmatter interface {
	MutableSubmat(Rect) Mutable
}

// Returns a thin wrapper which selects a constant submatrix.
// Calls A.ConstSubmat() if A is a ConstSubmatter.
func Submat(A Const, r Rect) Const {
	if A, ok := A.(ConstSubmatter); ok {
		return A.ConstSubmat(r)
	}
	return submatrixExpr{A, r}
}

type submatrixExpr struct {
	Matrix Const
	Rect   Rect
}

func (expr submatrixExpr) Size() Size { return expr.Rect.Size() }

func (expr submatrixExpr) At(i, j int) float64 {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	return expr.Matrix.At(i-i0, j-j0)
}

// Returns a thin wrapper which selects a mutable submatrix.
// Calls A.MutableSubmat() if A is a MutableSubmatter.
func MutableSubmat(A Mutable, r Rect) Mutable {
	if A, ok := A.(MutableSubmatter); ok {
		return A.MutableSubmat(r)
	}
	return mutableSubmatExpr{A, r}
}

type mutableSubmatExpr struct {
	Matrix Mutable
	Rect   Rect
}

func (expr mutableSubmatExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr mutableSubmatExpr) At(i, j int) float64 {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	return expr.Matrix.At(i-i0, j-j0)
}

func (expr mutableSubmatExpr) Set(i, j int, x float64) {
	i0 := expr.Rect.Min.I
	j0 := expr.Rect.Min.J
	expr.Matrix.Set(i-i0, j-j0, x)
}

type ConstReshaper interface {
	ConstReshape(Size) Const
}

type MutableReshaper interface {
	MutableReshape(Size) Mutable
}

// Address a constant matrix by a different shape.
// Equivalent to vectorizing then unvectorizing.
// If A is a ConstReshaper, calls ConstReshape().
func Reshape(A Const, rows, cols int) Const {
	if A, ok := A.(ConstReshaper); ok {
		return A.ConstReshape(Size{rows, cols})
	}
	return Unvec(Vec(A), rows, cols)
}

// Address a mutable matrix by a different shape.
// Equivalent to vectorizing then unvectorizing.
// If A is a MutableReshaper, calls MutableReshape().
func MutableReshape(A Mutable, rows, cols int) Mutable {
	if A, ok := A.(MutableReshaper); ok {
		return A.MutableReshape(Size{rows, cols})
	}
	return MutableUnvec(MutableVec(A), rows, cols)
}
