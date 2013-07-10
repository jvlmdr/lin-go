package mat

func Vec(A Const) VecExpr {
	return VecExpr{A}
}

// Vectorizes a constant matrix.
type VecExpr struct{ Const }

func (x VecExpr) Size() int {
	size := x.Const.Size()
	return size.Rows * size.Cols
}

func (x VecExpr) At(i int) float64 {
	M := x.Const
	rows := Rows(M)
	p := i / rows
	q := i % rows
	return M.At(p, q)
}

func MutableVec(A Mutable) MutableVecExpr {
	return MutableVecExpr{A}
}

// Vectorizes a mutable matrix.
type MutableVecExpr struct{ Mutable }

func (x MutableVecExpr) Size() int {
	size := x.Mutable.Size()
	return size.Rows * size.Cols
}

func (x MutableVecExpr) At(i int) float64 {
	M := x.Mutable
	rows := Rows(M)
	p := i / rows
	q := i % rows
	return M.At(p, q)
}

func (x MutableVecExpr) Set(i int, v float64) {
	M := x.Mutable
	rows := Rows(M)
	p := i / rows
	q := i % rows
	M.Set(p, q, v)
}

// Accesses a column in a constant matrix as a vector.
func Column(A Const, j int) ColumnExpr {
	return ColumnExpr{A, j}
}

// Accesses a column in a constant matrix as a vector.
type ColumnExpr struct {
	Matrix Const
	J      int
}

func (col ColumnExpr) Size() int {
	return Rows(col.Matrix)
}

func (col ColumnExpr) At(i int) float64 {
	return col.Matrix.At(i, col.J)
}

// Accesses a column in a mutable matrix as a vector.
func MutableColumn(A Mutable, j int) MutableColumnExpr {
	return MutableColumnExpr{A, j}
}

// Accesses a column in a mutable matrix as a vector.
type MutableColumnExpr struct {
	Matrix Mutable
	J      int
}

func (col MutableColumnExpr) Size() int {
	return Rows(col.Matrix)
}

func (col MutableColumnExpr) At(i int) float64 {
	return col.Matrix.At(i, col.J)
}

// Accesses a column in a constant matrix as a vector.
type RowExpr struct {
	Matrix Const
	I      int
}

func (row RowExpr) Size() int {
	return Rows(row.Matrix)
}

func (row RowExpr) At(j int) float64 {
	return row.Matrix.At(row.I, j)
}

// Accesses a column in a mutable matrix as a vector.
type MutableRowExpr struct {
	Matrix Mutable
	I      int
}

func (row MutableRowExpr) Size() int {
	return Rows(row.Matrix)
}

func (row MutableRowExpr) At(j int) float64 {
	return row.Matrix.At(row.I, j)
}

func (row MutableRowExpr) Set(j int, v float64) {
	row.Matrix.Set(row.I, j, v)
}
