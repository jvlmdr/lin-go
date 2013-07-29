package mat

import (
	"math"
	"testing"
)

func CheckEqual(t *testing.T, want, got Const, eps float64) {
	if !want.Size().Equals(got.Size()) {
		t.Fatalf("Matrix sizes do not match (want %v, got %v)", want.Size(), got.Size())
	}

	rows, cols := RowsCols(want)
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
