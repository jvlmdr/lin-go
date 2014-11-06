package blas_test

import (
	"math/rand"
	"testing"

	"github.com/jvlmdr/lin-go/blas"
	"github.com/jvlmdr/lin-go/mat"
)

func TestMatMul(t *testing.T) {
	const (
		m = 3
		k = 4
		n = 5
	)

	alpha := rand.NormFloat64()
	a, b := randMat(m, k), randMat(k, n)
	got := blas.MatMul(alpha, a, b)
	want := mat.Scale(alpha, mat.Mul(a, b))
	checkEqualMat(t, want, got, 1e-9)

	// Try with non-copying transposes.
	alpha = rand.NormFloat64()
	a, b = randMat(k, m).T(), randMat(k, n)
	got = blas.MatMul(alpha, a, b)
	want = mat.Scale(alpha, mat.Mul(a, b))
	checkEqualMat(t, want, got, 1e-9)

	alpha = rand.NormFloat64()
	a, b = randMat(m, k), randMat(n, k).T()
	got = blas.MatMul(alpha, a, b)
	want = mat.Scale(alpha, mat.Mul(a, b))
	checkEqualMat(t, want, got, 1e-9)

	alpha = rand.NormFloat64()
	a, b = randMat(k, m).T(), randMat(n, k).T()
	got = blas.MatMul(alpha, a, b)
	want = mat.Scale(alpha, mat.Mul(a, b))
	checkEqualMat(t, want, got, 1e-9)
}

func TestGenMatMul(t *testing.T) {
	const (
		m = 3
		k = 4
		n = 5
	)
	alpha, beta := rand.NormFloat64(), rand.NormFloat64()
	a, b, c := randMat(m, k), randMat(k, n), randMat(m, n)
	want := mat.Plus(mat.Scale(alpha, mat.Mul(a, b)), mat.Scale(beta, c))
	// Over-write c with result.
	blas.GenMatMul(alpha, a, b, beta, c)
	checkEqualMat(t, want, c, 1e-9)
}

func benchmarkMatMul(b *testing.B, m, k, n int, naive bool) {
	x, y := randMat(m, k), randMat(k, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if naive {
			mat.Mul(x, y)
		} else {
			blas.MatMul(1, x, y)
		}
	}
}

func BenchmarkMatMul_100x100_100x100(b *testing.B) {
	benchmarkMatMul(b, 100, 100, 100, false)
}

func BenchmarkMatMulNaive_100x100_100x100(b *testing.B) {
	benchmarkMatMul(b, 100, 100, 100, true)
}

func BenchmarkMatMul_1000x1000_1000x1000(b *testing.B) {
	benchmarkMatMul(b, 1000, 1000, 1000, false)
}

func BenchmarkMatMulNaive_1000x1000_1000x1000(b *testing.B) {
	benchmarkMatMul(b, 1000, 1000, 1000, true)
}
