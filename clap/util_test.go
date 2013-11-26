package clap

import (
	"github.com/jackvalmadre/lin-go/cmat"
	"math/cmplx"
	"math/rand"
	"testing"
)

const eps = 1e-12

func randNorm() complex128 {
	a := rand.NormFloat64()
	b := rand.NormFloat64()
	return complex(a, b)
}

func randMat(m, n int) *cmat.Mat {
	// Random symmetric positive definite matrix.
	a := cmat.New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a.Set(i, j, randNorm())
		}
	}
	return a
}

func randVec(n int) []complex128 {
	x := make([]complex128, n)
	for i := range x {
		x[i] = randNorm()
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

func epsEq(want, got complex128, eps float64) bool {
	return cmplx.Abs(want-got) <= eps
}

func testSliceEq(t *testing.T, want, got []complex128) {
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

func overDetProb(m, n int) (a *cmat.Mat, b, x []complex128, err error) {
	if m < n {
		panic("expect m >= n")
	}

	a = randMat(m, n)
	b = randVec(m)

	// Compute pseudo-inverse explicitly.
	// y <- (A' A) \ b
	x, err = SolveHerm(cmat.Mul(cmat.H(a), a), cmat.MulVec(cmat.H(a), b))
	if err != nil {
		return nil, nil, nil, err
	}
	return
}

func underDetProb(m, n int) (a *cmat.Mat, b, x []complex128, err error) {
	if m > n {
		panic("expect m <= n")
	}

	a = randMat(m, n)
	b = randVec(m)

	// Compute pseudo-inverse explicitly.
	// y <- (A A') \ b
	y, err := SolveHerm(cmat.Mul(a, cmat.H(a)), b)
	if err != nil {
		return nil, nil, nil, err
	}
	// x <- A' y
	x = cmat.MulVec(cmat.H(a), y)
	return
}
