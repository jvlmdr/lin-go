package mat

import "testing"

func TestMat(t *testing.T) {
	a := New(2, 3)
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 3)
	a.Set(1, 0, 4)
	a.Set(1, 1, 5)
	a.Set(1, 2, 6)

	ok := (a.At(0, 0) == 1 &&
		a.At(0, 1) == 2 &&
		a.At(0, 2) == 3 &&
		a.At(1, 0) == 4 &&
		a.At(1, 1) == 5 &&
		a.At(1, 2) == 6)

	if !ok {
		t.Fatalf("could not construct matrix and access elems")
	}
}

func TestNewRows(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	ok := (a.At(0, 0) == 1 &&
		a.At(0, 1) == 2 &&
		a.At(0, 2) == 3 &&
		a.At(1, 0) == 4 &&
		a.At(1, 1) == 5 &&
		a.At(1, 2) == 6)
	if !ok {
		t.Fatalf("could not construct matrix and access elems")
	}
}

func TestNewCols(t *testing.T) {
	got := NewCols([][]float64{
		{1, 4}, {2, 5}, {3, 6},
	})
	want := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	testMatEq(t, want, got)
}
