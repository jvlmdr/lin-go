package mat

// Returns a thin wrapper which transposes a constant matrix.
func T(A Const) Const {
	return tExpr{A}
}

type tExpr struct{ Matrix Const }

func (expr tExpr) Size() Size {
	return expr.Matrix.Size().T()
}

func (expr tExpr) At(i, j int) float64 {
	return expr.Matrix.At(j, i)
}

// Returns a thin wrapper which transposes a mutable matrix.
func MutableT(A Mutable) Mutable {
	return mutableTExpr{A}
}

type mutableTExpr struct{ Matrix Mutable }

func (expr mutableTExpr) Size() Size {
	return expr.Matrix.Size().T()
}

func (expr mutableTExpr) At(i, j int) float64 {
	return expr.Matrix.At(j, i)
}

func (expr mutableTExpr) Set(i, j int, v float64) {
	expr.Matrix.Set(j, i, v)
}

// Returns a thin wrapper which selects a constant submatrix.
func Submatrix(A Const, r Rect) Const {
	return submatrixExpr{A, r}
}

type submatrixExpr struct {
	Matrix Const
	Rect   Rect
}

func (expr submatrixExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr submatrixExpr) At(i, j int) float64 {
	return expr.Matrix.At(i-expr.Rect.Min.I, j-expr.Rect.Min.J)
}

// Returns a thin wrapper which selects a mutable submatrix.
func MutableSubmatrix(A Mutable, r Rect) Mutable {
	return mutableSubmatrixExpr{A, r}
}

type mutableSubmatrixExpr struct {
	Matrix Mutable
	Rect   Rect
}

func (expr mutableSubmatrixExpr) Size() Size {
	return expr.Rect.Size()
}

func (expr mutableSubmatrixExpr) At(i, j int) float64 {
	return expr.Matrix.At(i-expr.Rect.Min.I, j-expr.Rect.Min.J)
}

func (expr mutableSubmatrixExpr) Set(i, j int, x float64) {
	expr.Matrix.Set(i-expr.Rect.Min.I, j-expr.Rect.Min.J, x)
}
