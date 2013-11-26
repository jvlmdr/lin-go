package clap

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/cmat"
	"testing"
)

func TestLDLFact_Solve(t *testing.T) {
	n := 100
	// Random symmetric matrix.
	a := randMat(n, n)
	a = cmat.Plus(a, cmat.H(a))
	// Random vector.
	want := randVec(n)
	b := cmat.MulVec(a, want)

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
	a := cmat.NewRows([][]complex128{
		{1, 2},
		{2, -3},
	})
	// x = [1; 2]
	// b = A x = [5; -4]
	b := []complex128{5, -4}

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
	// [(1+0i) (2+0i)]
}
