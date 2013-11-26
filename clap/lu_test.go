package clap

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/cmat"
	"testing"
)

func TestLUFact_Solve(t *testing.T) {
	n := 100
	// Random square matrix.
	a := randMat(n, n)
	// Random vector.
	want := randVec(n)
	b := cmat.MulVec(a, want)

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
	b := cmat.MulVec(a, want)
	got, err := lu.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)

	// Then solve conjugate-transpose system.
	b = cmat.MulVec(cmat.H(a), want)
	got, err = lu.Solve(true, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleLUFact_Solve() {
	a := cmat.NewRows([][]complex128{
		{-1, 2},
		{3, 1},
	})
	// x = [1; 2]
	// b = A x = [3; 5]
	b := []complex128{3, 5}

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
	// [(1+0i) (2+0i)]
}
