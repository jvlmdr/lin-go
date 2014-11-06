package blas_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/jvlmdr/lin-go/blas"
	"github.com/jvlmdr/lin-go/mat"
)

func checkEqualMat(t *testing.T, want, got mat.Const, eps float64) bool {
	m, n := want.Dims()
	p, q := got.Dims()
	if !(m == p && n == q) {
		t.Errorf("different size: want %dx%d, got %dx%d", m, n, p, q)
		return false
	}
	ok := true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := want.At(i, j), got.At(i, j)
			if math.Abs(x-y) > eps {
				ok = false
				t.Errorf("different at %d, %d: want %g, got %g", i, j, x, y)
			}
		}
	}
	return ok
}

func randMat(m, n int) *blas.Mat {
	a := blas.NewMat(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a.Set(i, j, rand.NormFloat64())
		}
	}
	return a
}
