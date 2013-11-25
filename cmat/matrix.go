package cmat

type Mat struct {
	// Dimensions.
	Rows, Cols int
	// Elements in column-major order.
	// Element (i, j) resides at index (i + j*Rows).
	Elems []complex128
}

// Allocates a matrix of all zeros.
func New(m, n int) *Mat {
	elems := make([]complex128, m*n)
	return &Mat{m, n, elems}
}

func (a *Mat) Dims() (m, n int) {
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

// Creates a matrix from the list of rows.
//	NewRows([][]complex128{
//		{1, 2, 3},
//		{4, 5, 6},
//	})
func NewRows(rows [][]complex128) *Mat {
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
//	NewCols([][]complex128{{1, 4}, {2, 5}, {3, 6}})
func NewCols(cols [][]complex128) *Mat {
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

func eqLen(x [][]complex128) (n int, err error) {
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
