package clap

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/cmat"
	"testing"
)

func TestSolveSquare(t *testing.T) {
	n := 100
	a := randMat(n, n)
	want := randVec(n)
	b := cmat.MulVec(a, want)

	got, err := SolveSquare(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolveSquare() {
	a := cmat.NewRows([][]complex128{
		{-1, 2},
		{3, 1},
	})
	// x = [1; 2]
	// b = A x = [3; 5]
	b := []complex128{3, 5}

	x, err := SolveSquare(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [(1+0i) (2+0i)]
}
