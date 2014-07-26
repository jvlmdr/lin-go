package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestSolveSymm(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = mat.Mul(mat.T(a), a)

	// Random vector.
	want := randVec(n)
	b := mat.MulVec(a, want)

	// Factorize matrix.
	got, err := SolveSymm(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolveSymm() {
	a := mat.NewRows([][]float64{
		{1, 2},
		{2, -3},
	})
	// x = [1; 2]
	// b = A x = [5; -4]
	b := []float64{5, -4}

	x, err := SolveSymm(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}
