package lapack

import (
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zvec"
	"math"
	"math/cmplx"
	"testing"
)

func checkVectorsEqual(t *testing.T, want, got vec.Const, eps float64) {
	if want.Len() != got.Len() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
	}

	for i := 0; i < want.Len(); i++ {
		x := want.At(i)
		y := got.At(i)
		if math.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}

func checkComplexVectorsEqual(t *testing.T, want, got zvec.Const, eps float64) {
	if want.Len() != got.Len() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Len(), got.Len())
	}

	for i := 0; i < want.Len(); i++ {
		x := want.At(i)
		y := got.At(i)
		if cmplx.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}
