package cmat

import (
	"math/cmplx"
	"testing"
)

const eps = 1e-12

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
