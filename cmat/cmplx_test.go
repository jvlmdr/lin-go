package cmat

import "testing"

func TestH(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2 - 1i, 3 + 2i},
		{4 - 3i, 5 + 4i, -5i},
	})
	got := H(a)
	want := NewRows([][]complex128{
		{1, 4 + 3i},
		{2 + 1i, 5 - 4i},
		{3 - 2i, 5i},
	})
	testMatEq(t, want, got)
}

func TestConj(t *testing.T) {
	a := NewRows([][]complex128{
		{1, 2 - 1i, 3 + 2i},
		{4 - 3i, 5 + 4i, -5i},
	})
	got := Conj(a)
	want := NewRows([][]complex128{
		{1, 2 + 1i, 3 - 2i},
		{4 + 3i, 5 - 4i, 5i},
	})
	testMatEq(t, want, got)
}
