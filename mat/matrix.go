package mat

type Mat struct {
	// Dimensions.
	Rows, Cols int
	// Elements in column-major order.
	// Element (i, j) resides at index (i + j*Rows).
	Elems []float64
}

// Allocates a matrix of all zeros.
func New(m, n int) *Mat {
	elems := make([]float64, m*n)
	return &Mat{m, n, elems}
}

func (a *Mat) Dims() (m, n int) {
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

// Creates a matrix from the list of rows.
//	NewRows([][]float64{
//		{1, 2, 3},
//		{4, 5, 6},
//	})
func NewRows(rows [][]float64) *Mat {
	n, err := eqLen(rows)
	if err != nil {
		panic(err)
	}

	m := len(rows)
	a := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a.Set(i, j, rows[i][j])
		}
	}
	return a
}

// Creates a matrix from the list of columns.
//	NewCols([][]float64{{1, 4}, {2, 5}, {3, 6}})
func NewCols(cols [][]float64) *Mat {
	m, err := eqLen(cols)
	if err != nil {
		panic(err)
	}

	n := len(cols)
	a := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a.Set(i, j, cols[j][i])
		}
	}
	return a
}

func eqLen(x [][]float64) (n int, err error) {
	for i, xi := range x {
		if i == 0 {
			n = len(xi)
		} else {
			if len(xi) != n {
				return 0, errRagged(n, len(xi))
			}
		}
	}
	return
}
