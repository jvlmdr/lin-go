package lapack

import (
	"math"
	"math/rand"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

const eps = 1e-9

func randMat(m, n int) *mat.Mat {
	// Random symmetric positive definite matrix.
	a := mat.New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a.Set(i, j, rand.NormFloat64())
		}
	}
	return a
}

func randVec(n int) []float64 {
	x := make([]float64, n)
	for i := range x {
		x[i] = rand.NormFloat64()
	}
	return x
}

func testDimsEq(t *testing.T, want, got Const) {
	if !eqDims(want, got) {
		m, n := want.Dims()
		p, q := got.Dims()
		t.Fatalf("matrix sizes differ: want %dx%d, got %dx%d", m, n, p, q)
	}
}

func epsEq(want, got, eps float64) bool {
	return math.Abs(want-got) <= eps
}

func testSliceEq(t *testing.T, want, got []float64) {
	if len(want) != len(got) {
		t.Fatalf("lengths differ: want %d, got %d", len(want), len(got))
	}

	for i := range want {
		if !epsEq(want[i], got[i], eps) {
			t.Errorf("at %d: want %.4g, got %.4g", i, want[i], got[i])
		}
	}
}

func testMatEq(t *testing.T, want, got Const) {
	testDimsEq(t, want, got)

	m, n := want.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			u := want.At(i, j)
			v := got.At(i, j)
			if !epsEq(u, v, eps) {
				t.Errorf("at (%d, %d): want %.4g, got %.4g", i, j, u, v)
			}
		}
	}
}

func overDetProb(m, n int) (a *mat.Mat, b, x []float64, err error) {
	if m < n {
		panic("expect m >= n")
	}

	a = randMat(m, n)
	b = randVec(m)

	// Compute pseudo-inverse explicitly.
	// y <- (A' A) \ b
	x, err = SolveSymm(mat.Mul(mat.T(a), a), mat.MulVec(mat.T(a), b))
	if err != nil {
		return nil, nil, nil, err
	}
	return
}

func underDetProb(m, n int) (a *mat.Mat, b, x []float64, err error) {
	if m > n {
		panic("expect m <= n")
	}

	a = randMat(m, n)
	b = randVec(m)

	// Compute pseudo-inverse explicitly.
	// y <- (A A') \ b
	y, err := SolveSymm(mat.Mul(a, mat.T(a)), b)
	if err != nil {
		return nil, nil, nil, err
	}
	// x <- A' y
	x = mat.MulVec(mat.T(a), y)
	return
}
