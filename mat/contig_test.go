package mat

import "testing"

func TestContig_Size(t *testing.T) {
	s := MakeContig(2, 3).Size()
	if s.Rows != 2 {
		t.Errorf("Wrong number of rows")
	}
	if s.Cols != 3 {
		t.Errorf("Wrong number of rows")
	}
}

func TestContig_Set(t *testing.T) {
	A := MakeContig(2, 3)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	A.Set(0, 2, 5)
	A.Set(1, 2, 6)
}

func TestContig_At(t *testing.T) {
	A := MakeContig(2, 3)
	A.Set(1, 2, 42)
	if A.At(1, 2) != 42 {
		t.Errorf("At did not match Set")
	}
}

func TestContig_T(t *testing.T) {
	A := MakeContig(2, 3)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	A.Set(0, 2, 5)
	A.Set(1, 2, 6)
	B := A.T()

	want := Size{3, 2}
	got := B.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 1},
		{0, 1, 2},
		{1, 0, 3},
		{1, 1, 4},
		{2, 0, 5},
		{2, 1, 6},
	}

	for _, test := range tests {
		got := B.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, want)
		}
	}
}

func TestContig_Reshape(t *testing.T) {
	A := MakeContig(2, 3)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	A.Set(0, 2, 5)
	A.Set(1, 2, 6)
	B := A.Reshape(Size{3, 2})

	want := Size{3, 2}
	got := B.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 1},
		{1, 0, 2},
		{2, 0, 3},
		{0, 1, 4},
		{1, 1, 5},
		{2, 1, 6},
	}

	for _, test := range tests {
		got := B.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, want)
		}
	}
}

func TestContig_ColSlice_subset(t *testing.T) {
	A := MakeContig(2, 3)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	A.Set(0, 2, 5)
	A.Set(1, 2, 6)

	// Extract cols 1 and 2.
	B := A.ColSlice(1, 3)

	want := Size{2, 2}
	got := B.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 3},
		{1, 0, 4},
		{0, 1, 5},
		{1, 1, 6},
	}

	for _, test := range tests {
		got := B.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, test.want)
		}
	}
}

func TestContig_ColSlice_grow(t *testing.T) {
	A := MakeContigCap(2, 3, 6)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	A.Set(0, 2, 5)
	A.Set(1, 2, 6)

	// Extract cols 2 and 3.
	B := A.ColSlice(2, 4)

	want := Size{2, 2}
	got := B.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 5},
		{1, 0, 6},
		{0, 1, 0},
		{1, 1, 0},
	}

	for _, test := range tests {
		got := B.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, test.want)
		}
	}
}

func TestContig_ColAppend_grow(t *testing.T) {
	A := MakeContig(2, 2)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	B := MakeContig(2, 2)
	B.Set(0, 0, 5)
	B.Set(1, 0, 6)
	B.Set(0, 1, 7)
	B.Set(1, 1, 8)
	// Append A and B.
	C := A.ColAppend(B)

	want := Size{2, 4}
	got := C.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 1},
		{1, 0, 2},
		{0, 1, 3},
		{1, 1, 4},
		{0, 2, 5},
		{1, 2, 6},
		{0, 3, 7},
		{1, 3, 8},
	}
	for _, test := range tests {
		got := C.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, test.want)
		}
	}
}

func TestContig_ColAppend_inPlace(t *testing.T) {
	A := MakeContigCap(2, 2, 4)
	A.Set(0, 0, 1)
	A.Set(1, 0, 2)
	A.Set(0, 1, 3)
	A.Set(1, 1, 4)
	B := MakeContig(2, 2)
	B.Set(0, 0, 5)
	B.Set(1, 0, 6)
	B.Set(0, 1, 7)
	B.Set(1, 1, 8)
	// Append A and B.
	C := A.ColAppend(B)

	want := Size{2, 4}
	got := C.Size()
	if !got.Equals(want) {
		t.Fatalf("Wrong size (want %v, got %v)", want, got)
	}

	var tests = []struct {
		i    int
		j    int
		want float64
	}{
		{0, 0, 1},
		{1, 0, 2},
		{0, 1, 3},
		{1, 1, 4},
		{0, 2, 5},
		{1, 2, 6},
		{0, 3, 7},
		{1, 3, 8},
	}
	for _, test := range tests {
		got := C.At(test.i, test.j)
		if got != test.want {
			t.Errorf("Wrong value at (%d, %d) (want %g, got %g)",
				test.i, test.j, got, test.want)
		}
	}

	if &A.Elems[0] != &C.Elems[0] {
		t.Fatalf("Append was not in place")
	}
}
