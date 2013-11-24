package mat

import "testing"

func TestT(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := T(a)
	want := NewRows([][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	})
	testMatEq(t, want, got)
}

func TestDiag(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2},
		{4, 5},
	})
	got := Diag(a)
	want := []float64{1, 5}
	testSliceEq(t, want, got)
}

func TestDiag_fat(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Diag(a)
	want := []float64{1, 5}
	testSliceEq(t, want, got)
}

func TestDiag_skinny(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	got := Diag(a)
	want := []float64{1, 4}
	testSliceEq(t, want, got)
}

func TestNewDiag(t *testing.T) {
	got := NewDiag([]float64{-1, 2, 1})
	want := NewRows([][]float64{
		{-1, 0, 0},
		{0, 2, 0},
		{0, 0, 1},
	})
	testMatEq(t, want, got)
}

func TestI(t *testing.T) {
	got := I(3)
	want := NewRows([][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	testMatEq(t, want, got)
}

func TestAugment(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]float64{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Augment(a, b)
	want := NewRows([][]float64{
		{1, 2, 3, -1, 0, 1},
		{4, 5, 6, -1, 2, -1},
	})
	testMatEq(t, want, got)
}

func TestStack(t *testing.T) {
	a := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]float64{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Stack(a, b)
	want := NewRows([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{-1, 0, 1},
		{-1, 2, -1},
	})
	testMatEq(t, want, got)
}
