package blas

import "fmt"

// MatMul computes the product of two matrices.
func MatMul(alpha float64, a, b *Mat) *Mat {
	c := NewMat(a.Rows, b.Cols)
	GenMatMul(alpha, a, b, 0, c)
	return c
}

// GenMatMul computes alpha a b + beta c and stores the result in c.
// The destination c cannot be row-major.
func GenMatMul(alpha float64, a, b *Mat, beta float64, c *Mat) {
	if err := errIfNotCompat(a, b); err != nil {
		panic(err)
	}
	if err := errIfDimsNotEq(a.Rows, b.Cols, c.Rows, c.Cols); err != nil {
		panic(err)
	}
	if err := errIfDstInvalid(c); err != nil {
		panic(err)
	}
	ta := trans(a)
	tb := trans(b)
	m, k, n := a.Rows, a.Cols, b.Cols
	dgemm(ta, tb, m, n, k, alpha, a.Elems, a.Stride, b.Elems, b.Stride, beta, c.Elems, c.Stride)
}

func errIfNotCompat(a, b *Mat) error {
	if a.Cols != b.Rows {
		return fmt.Errorf("incompatible dims: %dx%d, %dx%d", a.Rows, a.Cols, b.Rows, b.Cols)
	}
	return nil
}

func errIfDimsNotEq(m, n, p, q int) error {
	if !(m == p && n == q) {
		return fmt.Errorf("different dims: %dx%d, %dx%d", m, n, p, q)
	}
	return nil
}

func errIfDstInvalid(a *Mat) error {
	if a.RowMaj {
		return fmt.Errorf("destination matrix is row major")
	}
	return nil
}
