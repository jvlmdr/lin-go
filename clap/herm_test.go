package clap

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/cmat"
	"testing"
)

func TestSolveHerm(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = cmat.Mul(cmat.H(a), a)

	// Random vector.
	want := randVec(n)
	b := cmat.MulVec(a, want)

	// Factorize matrix.
	got, err := SolveHerm(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolveHerm() {
	a := cmat.NewRows([][]complex128{
		{1, 2},
		{2, -3},
	})
	// x = [1; 2]
	// b = A x = [5; -4]
	b := []complex128{5, -4}

	x, err := SolveHerm(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x, 'f', 3))
	// Output:
	// [(1.000+0.000i) (2.000+0.000i)]
}
