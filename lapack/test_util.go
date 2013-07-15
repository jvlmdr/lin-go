package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zvec"
	"math"
	"math/cmplx"
	"testing"
)

func checkVectorsEqual(t *testing.T, want, got vec.Const, eps float64) {
	if want.Size() != got.Size() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Size(), got.Size())
	}

	for i := 0; i < want.Size(); i++ {
		x := want.At(i)
		y := got.At(i)
		if math.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}

func checkMatricesEqual(t *testing.T, want, got mat.Const, eps float64) {
	if !want.Size().Equals(got.Size()) {
		t.Fatalf("Matrix sizes do not match (want %v, got %v)", want.Size(), got.Size())
	}

	rows, cols := mat.RowsCols(want)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			x := want.At(i, j)
			y := got.At(i, j)
			if math.Abs(x-y) > 1e-6 {
				t.Errorf("Wrong value at %d, %d (want %g, got %g)", i, j, x, y)
			}
		}
	}
}

func checkComplexVectorsEqual(t *testing.T, want, got zvec.Const, eps float64) {
	if want.Size() != got.Size() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Size(), got.Size())
	}

	for i := 0; i < want.Size(); i++ {
		x := want.At(i)
		y := got.At(i)
		if cmplx.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}
