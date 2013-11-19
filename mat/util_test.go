package mat

import (
	"math"
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

//	func testEqualMatCmplx(t *testing.T, want, got zmat.Const, eps float64) {
//		if !want.Size().Equals(got.Size()) {
//			t.Fatalf("Matrix sizes do not match (want %v, got %v)", want.Size(), got.Size())
//		}
//	
//		rows, cols := zmat.RowsCols(want)
//		for i := 0; i < rows; i++ {
//			for j := 0; j < cols; j++ {
//				x := want.At(i, j)
//				y := got.At(i, j)
//				if cmplx.Abs(x-y) > eps {
//					t.Errorf("Wrong value at %d, %d (want %g, got %g)", i, j, x, y)
//				}
//			}
//		}
//	}
//	
//	func testEqualVec(t *testing.T, want, got vec.Const, eps float64) {
//		if want.Len() != got.Len() {
//			t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
//		}
//	
//		for i := 0; i < want.Len(); i++ {
//			x := want.At(i)
//			y := got.At(i)
//			if math.Abs(x-y) > eps {
//				t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
//			}
//		}
//	}
//	
//	func testEqualVecCmplx(t *testing.T, want, got zvec.Const, eps float64) {
//		if want.Len() != got.Len() {
//			t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
//		}
//	
//		for i := 0; i < want.Len(); i++ {
//			x := want.At(i)
//			y := got.At(i)
//			if cmplx.Abs(x-y) > eps {
//				t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
//			}
//		}
//	}
