package vec

// Lazy evaluators for vector-valued functions.
// Idiomatic use is such that x <- x - y is written vec.Copy(x, Minus(x, y)).

// Addition of two vectors.
// Lazily evaluated.
func Plus(x, y Const) PlusExpr { return PlusExpr{x, y} }

type PlusExpr struct{ X, Y Const }

func (expr PlusExpr) Size() int        { return expr.X.Size() }
func (expr PlusExpr) At(i int) float64 { return expr.X.At(i) + expr.Y.At(i) }

// Difference between two vectors.
// Lazily evaluated.
func Minus(x, y Const) MinusExpr { return MinusExpr{x, y} }

type MinusExpr struct{ X, Y Const }

func (expr MinusExpr) Size() int        { return expr.X.Size() }
func (expr MinusExpr) At(i int) float64 { return expr.X.At(i) - expr.Y.At(i) }

// Scalar multiplication.
// Lazily evaluated.
func Scale(alpha float64, x Const) ScaleExpr { return ScaleExpr{alpha, x} }

type ScaleExpr struct {
	Alpha float64
	X     Const
}

func (expr ScaleExpr) Size() int        { return expr.X.Size() }
func (expr ScaleExpr) At(i int) float64 { return expr.Alpha * expr.X.At(i) }

// Constant vector of ones.
func One(n int) OneExpr { return OneExpr{n} }

type OneExpr struct{ N int }

func (expr OneExpr) Size() int      { return expr.N }
func (expr OneExpr) At(int) float64 { return 1 }

// Constant vector of zeros.
func Zero(n int) ZeroExpr { return ZeroExpr{n} }

type ZeroExpr struct{ N int }

func (expr ZeroExpr) Size() int      { return expr.N }
func (expr ZeroExpr) At(int) float64 { return 0 }

// Constant vector.
func Constant(n int, x float64) ConstantExpr { return ConstantExpr{n, x} }

type ConstantExpr struct {
	N int
	X float64
}

func (expr ConstantExpr) Size() int      { return expr.N }
func (expr ConstantExpr) At(int) float64 { return expr.X }

// Applies the same unary function to every element.
// Lazily evaluated.
func Map(x Const, f func(int, float64) float64) MapExpr { return MapExpr{x, f} }

type MapExpr struct {
	X Const
	F func(int, float64) float64
}

func (expr MapExpr) Size() int        { return expr.X.Size() }
func (expr MapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i)) }

// Applies the same binary function to every element.
// Lazily evaluated.
func BinaryMap(x, y Const, f func(int, float64, float64) float64) BinaryMapExpr {
	return BinaryMapExpr{x, y, f}
}

type BinaryMapExpr struct {
	X Const
	Y Const
	F func(int, float64, float64) float64
}

func (expr BinaryMapExpr) Size() int        { return expr.X.Size() }
func (expr BinaryMapExpr) At(i int) float64 { return expr.F(i, expr.X.At(i), expr.Y.At(i)) }
