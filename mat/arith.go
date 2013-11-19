package mat

// Adds two matrices.
//
// Panics if matrices have different dimensions.
func Plus(a, b Const) Mutable {
	if err := errIfDimsNotEq(a, b); err != nil {
		panic(err)
	}

	m, n := a.Dims()
	c := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c.Set(i, j, a.At(i, j)+b.At(i, j))
		}
	}
	return c
}

// Subtracts one matrix from another.
//
// Panics if matrices have different dimensions.
func Minus(a, b Const) Mutable {
	if err := errIfDimsNotEq(a, b); err != nil {
		panic(err)
	}

	m, n := a.Dims()
	c := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c.Set(i, j, a.At(i, j)-b.At(i, j))
		}
	}
	return c
}

// Scales a matrix by a constant.
func Scale(k float64, a Const) Mutable {
	m, n := a.Dims()
	b := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b.Set(i, j, k*a.At(i, j))
		}
	}
	return b
}

// Multiplies two matrices.
//
// Panics if the dimensions are incompatible.
func Mul(a, b Const) Mutable {
	if err := errMulDimsMat(a, b); err != nil {
		panic(err)
	}

	m, n := a.Dims()
	_, q := b.Dims()
	c := New(m, q)
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			var v float64
			for k := 0; k < n; k++ {
				v += a.At(i, k) * b.At(k, j)
			}
			c.Set(i, j, v)
		}
	}
	return c
}

// Computes y = A x.
//
// Panics if the dimensions are incompatible.
func MulVec(a Const, x []float64) []float64 {
	if err := errMulDimsVec(a, x); err != nil {
		panic(err)
	}

	m, n := a.Dims()
	y := make([]float64, m)
	for i := 0; i < m; i++ {
		var v float64
		for k := 0; k < n; k++ {
			v += a.At(i, k) * x[k]
		}
		y[i] = v
	}
	return y
}
