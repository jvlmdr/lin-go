package clap

type Const interface {
	// Returns matrix size.
	Dims() (rows, cols int)
	// Accesses element.
	At(i, j int) complex128
}

type Mat struct {
	// Dimensions.
	Rows, Cols int
	// Elements in column-major order.
	// Element (i, j) resides at index (i + j*Rows).
	Elems []complex128
}

// Allocates a matrix of all zeros.
func NewMat(rows, cols int) *Mat {
	elems := make([]complex128, rows*cols)
	return &Mat{rows, cols, elems}
}

func (a *Mat) Dims() (rows, cols int) {
	return a.Rows, a.Cols
}

func (a *Mat) At(i, j int) complex128 {
	return a.Elems[a.index(i, j)]
}

func (a *Mat) Set(i, j int, v complex128) {
	a.Elems[a.index(i, j)] = v
}

// Returns the column-major index of element (i, j).
func (a *Mat) index(i, j int) int {
	return i + j*a.Rows
}
