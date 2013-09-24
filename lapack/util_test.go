package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"math"
	"math/cmplx"
	"testing"
)

func checkEqualMat(t *testing.T, want, got mat.Const, eps float64) {
	if !want.Size().Equals(got.Size()) {
		t.Fatalf("Matrix sizes do not match (want %v, got %v)", want.Size(), got.Size())
	}

	rows, cols := mat.RowsCols(want)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := want.At(i, j)
			y := got.At(i, j)
			if math.Abs(x-y) > eps {
				t.Errorf("Wrong value at %d, %d (want %g, got %g)", i, j, x, y)
			}
		}
	}
}

func checkEqualMatCmplx(t *testing.T, want, got zmat.Const, eps float64) {
	if !want.Size().Equals(got.Size()) {
		t.Fatalf("Matrix sizes do not match (want %v, got %v)", want.Size(), got.Size())
	}

	rows, cols := zmat.RowsCols(want)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := want.At(i, j)
			y := got.At(i, j)
			if cmplx.Abs(x-y) > eps {
				t.Errorf("Wrong value at %d, %d (want %g, got %g)", i, j, x, y)
			}
		}
	}
}

func checkEqualVec(t *testing.T, want, got vec.Const, eps float64) {
	if want.Len() != got.Len() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
	}

	for i := 0; i < want.Len(); i++ {
		x := want.At(i)
		y := got.At(i)
		if math.Abs(x-y) > eps {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}

func checkEqualVecCmplx(t *testing.T, want, got zvec.Const, eps float64) {
	if want.Len() != got.Len() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
	}

	for i := 0; i < want.Len(); i++ {
		x := want.At(i)
		y := got.At(i)
		if cmplx.Abs(x-y) > eps {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}
