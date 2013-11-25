package lapack

type Const interface {
	// Returns matrix size.
	Dims() (rows, cols int)
	// Accesses element.
	At(i, j int) float64
}

type Mat struct {
	// Dimensions.
	Rows, Cols int
	// Elements in column-major order.
	// Element (i, j) resides at index (i + j*Rows).
	Elems []float64
}

// Allocates a matrix of all zeros.
func NewMat(rows, cols int) *Mat {
	elems := make([]float64, rows*cols)
	return &Mat{rows, cols, elems}
}

func (a *Mat) Dims() (rows, cols int) {
	return a.Rows, a.Cols
}

func (a *Mat) At(i, j int) float64 {
	return a.Elems[a.index(i, j)]
}

func (a *Mat) Set(i, j int, v float64) {
	a.Elems[a.index(i, j)] = v
}

// Returns the column-major index of element (i, j).
func (a *Mat) index(i, j int) int {
	return i + j*a.Rows
}
