package cmat

import "testing"

func TestT(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := T(a)
	want := NewRows([][]complex128{
		{1, 4},
		{2, 5},
		{3, 6},
	})
	testMatEq(t, want, got)
}

func TestDiag(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2},
		{4, 5},
	})
	got := Diag(a)
	want := []complex128{1, 5}
	testSliceEq(t, want, got)
}

func TestDiag_fat(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Diag(a)
	want := []complex128{1, 5}
	testSliceEq(t, want, got)
}

func TestDiag_skinny(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	got := Diag(a)
	want := []complex128{1, 4}
	testSliceEq(t, want, got)
}

func TestNewDiag(t *testing.T) {
	got := NewDiag([]complex128{-1, 2, 1})
	want := NewRows([][]complex128{
		{-1, 0, 0},
		{0, 2, 0},
		{0, 0, 1},
	})
	testMatEq(t, want, got)
}

func TestI(t *testing.T) {
	got := I(3)
	want := NewRows([][]complex128{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
	testMatEq(t, want, got)
}

func TestVec(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Vec(a)
	want := []complex128{1, 4, 2, 5, 3, 6}
	testSliceEq(t, want, got)
}

func TestUnvec(t *testing.T) {
	x := []complex128{1, 2, 3, 4, 5, 6}
	got := Unvec(x, 2, 3)
	want := NewRows([][]complex128{
		{1, 3, 5},
		{2, 4, 6},
	})
	testMatEq(t, want, got)
}

func TestRow(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Row(a, 1)
	want := []complex128{4, 5, 6}
	testSliceEq(t, want, got)
}

func TestCol(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Col(a, 1)
	want := []complex128{2, 5}
	testSliceEq(t, want, got)
}

func TestAugment(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Augment(a, b)
	want := NewRows([][]complex128{
		{1, 2, 3, -1, 0, 1},
		{4, 5, 6, -1, 2, -1},
	})
	testMatEq(t, want, got)
}

func TestStack(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Stack(a, b)
	want := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
		{-1, 0, 1},
		{-1, 2, -1},
	})
	testMatEq(t, want, got)
}
