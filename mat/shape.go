package mat

// Copies the matrix into a transposed matrix.
func T(a Const) Mutable {
	m, n := a.Dims()
	at := New(n, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			at.Set(j, i, a.At(i, j))
		}
	}
	return at
}

// Extracts the diagonal of a matrix.
func Diag(a Const) []float64 {
	m, n := a.Dims()
	d := make([]float64, min(m, n))
	for i := range d {
		d[i] = a.At(i, i)
	}
	return d
}

// Constructs a square, diagonal matrix.
func NewDiag(d []float64) Mutable {
	n := len(d)
	a := New(n, n)
	for i, v := range d {
		a.Set(i, i, v)
	}
	return a
}

// Instantiates an identity matrix.
func I(n int) Mutable {
	a := New(n, n)
	for i := 0; i < n; i++ {
		a.Set(i, i, 1)
	}
	return a
}

// Returns the horizontal concatenation [a_{0}, a_{1}, ..., a_{n-1}].
//
// Panics if matrices have different numbers of rows.
func Augment(srcs ...Const) Mutable {
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

// Returns the vertical concatenation [a_{0}; a_{1}; ...; a_{n-1}].
//
// Panics if matrices have different numbers of columns.
func Stack(srcs ...Const) Mutable {
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
