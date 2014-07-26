package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestLUFact_Solve(t *testing.T) {
	n := 100
	// Random square matrix.
	a := randMat(n, n)
	// Random vector.
	want := randVec(n)
	b := mat.MulVec(a, want)

	lu, err := LU(a)
	if err != nil {
		t.Fatal(err)
	}
	got, err := lu.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func TestLUFact_Solve_t(t *testing.T) {
	n := 100
	// Random square matrix.
	a := randMat(n, n)
	// Random vector.
	want := randVec(n)

	// Factorize.
	lu, err := LU(a)
	if err != nil {
		t.Fatal(err)
	}

	// Solve un-transposed system.
	b := mat.MulVec(a, want)
	got, err := lu.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)

	// Then solve transposed system.
	b = mat.MulVec(mat.T(a), want)
	got, err = lu.Solve(true, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleLUFact_Solve() {
	a := mat.NewRows([][]float64{
		{-1, 2},
		{3, 1},
	})
	// x = [1; 2]
	// b = A x = [3; 5]
	b := []float64{3, 5}

	lu, err := LU(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := lu.Solve(false, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}
