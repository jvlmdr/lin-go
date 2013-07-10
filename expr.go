package mat

// Returns a thin wrapper which transposes a constant matrix.
func T(A Const) TExpr {
	return TExpr{A}
}

// Thin wrapper which transposes a constant matrix.
type TExpr struct {
	Matrix Const
}

func (expr TExpr) Size() Size {
	return expr.Matrix.Size().T()
}

func (expr TExpr) At(i, j int) float64 {
	return expr.Matrix.At(j, i)
}

// Returns a thin wrapper which transposes a mutable matrix.
func MutableT(A Const) MutableTExpr {
	return MutableTExpr{A, r}
}

// Thin wrapper which transposes a mutable matrix.
type MutableTExpr struct {
	Matrix Const
	Rect   Rectangle
}

func (expr MutableTExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr MutableTExpr) At(i, j int) float64 {
	return expr.Matrix.At(i-expr.Rect.Min.I, j-expr.Rect.Min.J)
}

// Returns a thin wrapper which selects a constant submatrix.
func Submatrix(A Const, r Rectangle) ConstSubmatrixExpr {
	return ConstSubmatrixExpr{A, r}
}

// Thin wrapper which selects a constant submatrix.
type ConstSubmatrixExpr struct {
	Matrix Const
	Rect   Rectangle
}

func (expr ConstSubmatrixExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr ConstSubmatrixExpr) At(i, j int) float64 {
	return expr.Matrix.At(i-expr.Rect.Min.I, j-expr.Rect.Min.J)
}

// Returns a thin wrapper which selects a mutable submatrix.
func MutableSubmatrix(A Mutable, r Rectangle) MutableSubmatrixExpr {
	return MutableSubmatrixExpr{A, r}
}

// Thin wrapper which selects a mutable submatrix.
type MutableSubmatrixExpr struct {
	Matrix Mutable
	Rect   Rectangle
}

func (expr MutableSubmatrixExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr MutableSubmatrixExpr) At(i, j int) float64 {
	return expr.Matrix.At(i-expr.Rect.Min.I, j-expr.Rect.Min.J)
}

func (expr MutableSubmatrixExpr) Set(i, j int, x float64) {
	expr.Matrix.Set(i-expr.Rect.Min.I, j-expr.Rect.Min.J, x)
}
