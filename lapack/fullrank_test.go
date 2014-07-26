package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestSolveFullRank_overdetermined(t *testing.T) {
	m, n := 150, 100
	a, b, want, err := overDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := SolveFullRank(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func TestSolveFullRank_underdetermined(t *testing.T) {
	m, n := 100, 150
	a, b, want, err := underDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := SolveFullRank(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolveFullRank_overdetermined() {
	// Find minimum-error solution to
	//	x     = 3,
	//	    y = 6,
	//	x + y = 3.
	a := mat.NewRows([][]float64{
		{1, 0},
		{0, 1},
		{1, 1},
	})
	b := []float64{3, 6, 3}

	x, err := SolveFullRank(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 4]
}

func ExampleSolveFullRank_underdetermined() {
	// Find minimum-norm solution to
	//	x     + z = 6,
	//	    y + z = 9.
	a := mat.NewRows([][]float64{
		{1, 0, 1},
		{0, 1, 1},
	})
	b := []float64{6, 9}

	x, err := SolveFullRank(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 4 5]
}
