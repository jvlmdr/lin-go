package vec

// Vector whose i-th element is f(x.At(i)).
func Map(x Const, f func(float64) float64) Const { return mapExpr{x, f} }

type mapExpr struct {
	X Const
	F func(float64) float64
}

func (expr mapExpr) Size() int        { return expr.X.Size() }
func (expr mapExpr) At(i int) float64 { return expr.F(expr.X.At(i)) }

// Vector whose i-th element is f(x.At(i), y.At(i)).
func MapTwo(x, y Const, f func(float64, float64) float64) Const {
	return mapTwoExpr{x, y, f}
}

type mapTwoExpr struct {
	X Const
	Y Const
	F func(float64, float64) float64
}

func (expr mapTwoExpr) Size() int        { return expr.X.Size() }
func (expr mapTwoExpr) At(i int) float64 { return expr.F(expr.X.At(i), expr.Y.At(i)) }

// Vector whose i-th element is f().
func MapNil(n int, f func() float64) Const { return mapNilExpr{n, f} }

type mapNilExpr struct {
	N int
	F func() float64
}

func (expr mapNilExpr) Size() int        { return expr.N }
func (expr mapNilExpr) At(i int) float64 { return expr.F() }

// Vector whose i-th element is f(i).
func MapIndex(n int, f func(int) float64) Const { return mapIndexExpr{n, f} }

type mapIndexExpr struct {
	N int
	F func(int) float64
}

func (expr mapIndexExpr) Size() int        { return expr.N }
func (expr mapIndexExpr) At(i int) float64 { return expr.F(i) }
