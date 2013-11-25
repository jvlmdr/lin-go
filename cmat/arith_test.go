package cmat

import "testing"

func TestPlus(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Plus(a, b)
	want := NewRows([][]complex128{
		{0, 2, 4},
		{3, 7, 5},
	})
	testMatEq(t, want, got)
}

func TestMinus(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	b := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Minus(a, b)
	want := NewRows([][]complex128{
		{2, 2, 2},
		{5, 3, 7},
	})
	testMatEq(t, want, got)
}

func TestScale(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2, 3},
		{4, 5, 6},
	})
	got := Scale(-2, a)
	want := NewRows([][]complex128{
		{-2, -4, -6},
		{-8, -10, -12},
	})
	testMatEq(t, want, got)
}

func TestMul(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2},
		{4, 5},
	})
	b := NewRows([][]complex128{
		{-1, 0, 1},
		{-1, 2, -1},
	})
	got := Mul(a, b)
	want := NewRows([][]complex128{
		{-3, 4, -1},
		{-9, 10, -1},
	})
	testMatEq(t, want, got)
}

func TestMulVec(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	b := []complex128{-2, 1}
	got := MulVec(a, b)
	want := []complex128{0, -2, -4}
	testSliceEq(t, want, got)
}
