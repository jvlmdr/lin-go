package mat

// Creates a transposed copy of the matrix.
func T(a Const) *Mat {
	m, n := a.Dims()
	at := New(n, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			at.Set(j, i, a.At(i, j))
		}
	}
	return at
}

// Instantiates an identity matrix.
func I(n int) *Mat {
	a := New(n, n)
	for i := 0; i < n; i++ {
		a.Set(i, i, 1)
	}
	return a
}

// Returns the concatenation of the columns of a matrix.
func Vec(a Const) []float64 {
	m, n := a.Dims()
	x := make([]float64, 0, m*n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			x = append(x, a.At(i, j))
		}
	}
	return x
}

// Reconstructs a matrix from a vectorized matrix.
func Unvec(x []float64, m, n int) *Mat {
	a := New(m, n)
	var k int
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			a.Set(i, j, x[k])
			k++
		}
	}
	return a
}

// Returns a copy of the elements of a row.
func Row(a Const, i int) []float64 {
	_, n := a.Dims()
	x := make([]float64, n)
	for j := 0; j < n; j++ {
		x[j] = a.At(i, j)
	}
	return x
}

// Returns a copy of the elements of a column.
func Col(a Const, j int) []float64 {
	m, _ := a.Dims()
	x := make([]float64, m)
	for i := 0; i < m; i++ {
		x[i] = a.At(i, j)
	}
	return x
}

// Copies a row into a matrix.
func SetRow(a Mutable, i int, x []float64) {
	_, n := a.Dims()
	for j := 0; j < n; j++ {
		a.Set(i, j, x[j])
	}
}

// Copies a column into a matrix.
func SetCol(a Mutable, j int, x []float64) {
	m, _ := a.Dims()
	for i := 0; i < m; i++ {
		a.Set(i, j, x[i])
	}
}

// Extracts the diagonal of the matrix.
func Diag(a Const) []float64 {
	m, n := a.Dims()
	d := make([]float64, min(m, n))
	for i := range d {
		d[i] = a.At(i, i)
	}
	return d
}

// Modifies the diagonal of the matrix.
func SetDiag(a Mutable, d []float64) {
	for i := range d {
		a.Set(i, i, d[i])
	}
}

// Constructs a square, diagonal matrix.
func NewDiag(d []float64) *Mat {
	n := len(d)
	a := New(n, n)
	for i, v := range d {
		a.Set(i, i, v)
	}
	return a
}

// Returns the horizontal concatenation of the matrices.
//
// Panics if matrices have different numbers of rows.
func Augment(srcs ...Const) *Mat {
	rows, cols := augmentDims(srcs...)
	dst := New(rows, cols)
	var off int
	for _, a := range srcs {
		m, n := a.Dims()
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dst.Set(i, j+off, a.At(i, j))
			}
		}
		off += n
	}
	return dst
}

// Panics if matrices have different numbers of rows.
func augmentDims(srcs ...Const) (rows, cols int) {
	for i, a := range srcs {
		m, n := a.Dims()
		if i == 0 {
			rows = m
		} else {
			if m != rows {
				panic(errRagged(rows, m))
			}
		}
		cols += n
	}
	return
}

// Returns the vertical concatenation of the matrices.
//
// Panics if matrices have different numbers of columns.
func Stack(srcs ...Const) *Mat {
	rows, cols := stackDims(srcs...)
	dst := New(rows, cols)
	var off int
	for _, a := range srcs {
		m, n := a.Dims()
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dst.Set(i+off, j, a.At(i, j))
			}
		}
		off += m
	}
	return dst
}

// Panics if matrices have different numbers of columns.
func stackDims(srcs ...Const) (rows, cols int) {
	for i, a := range srcs {
		m, n := a.Dims()
		if i == 0 {
			cols = n
		} else {
			if n != cols {
				panic(errRagged(cols, n))
			}
		}
		rows += m
	}
	return
}
