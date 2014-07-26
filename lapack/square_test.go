package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestSolveSquare(t *testing.T) {
	n := 100
	a := randMat(n, n)
	want := randVec(n)
	b := mat.MulVec(a, want)

	got, err := SolveSquare(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolveSquare() {
	a := mat.NewRows([][]float64{
		{-1, 2},
		{3, 1},
	})
	// x = [1; 2]
	// b = A x = [3; 5]
	b := []float64{3, 5}

	x, err := SolveSquare(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}
