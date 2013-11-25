package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"testing"
)

func TestLDLFact_Solve(t *testing.T) {
	n := 100
	// Random symmetric matrix.
	a := randMat(n, n)
	a = mat.Plus(a, mat.T(a))
	// Random vector.
	want := randVec(n)
	b := mat.MulVec(a, want)

	ldl, err := LDL(a)
	if err != nil {
		t.Fatal(err)
	}

	got, err := ldl.Solve(b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleLDLFact_Solve() {
	a := mat.NewRows([][]float64{
		{1, 2},
		{2, -3},
	})
	// x = [1; 2]
	// b = A x = [5; -4]
	b := []float64{5, -4}

	ldl, err := LDL(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := ldl.Solve(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}
