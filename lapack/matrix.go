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

func (A *Mat) Dims() (rows, cols int) {
	return A.Rows, A.Cols
}

func (A *Mat) At(i, j int) float64 {
	return A.Elems[A.index(i, j)]
}

func (A *Mat) Set(i, j int, v float64) {
	A.Elems[A.index(i, j)] = v
}

// Returns the column-major index of element (i, j).
func (A *Mat) index(i, j int) int {
	return i + j*A.Rows
}

func transpose(src Const) *Mat {
	m, n := src.Dims()
	dst := NewMat(n, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dst.Set(j, i, src.At(i, j))
		}
	}
	return dst
}
