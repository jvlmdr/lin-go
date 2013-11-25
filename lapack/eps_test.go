package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"testing"
)

func TestSolve_overdetermined(t *testing.T) {
	m, n := 150, 100
	a, b, want, err := overDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := Solve(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func TestSolve_underdetermined(t *testing.T) {
	m, n := 100, 150
	a, b, want, err := underDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	got, err := Solve(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolve_overdetermined() {
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

	x, err := Solve(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g\n", x)
	// Output:
	// [1 4]
}

// Find minimum-norm solution to
//	x     + z = 6,
//	    y + z = 9.
func ExampleSolve_underdetermined() {
	a := mat.NewRows([][]float64{
		{1, 0, 1},
		{0, 1, 1},
	})
	b := []float64{6, 9}

	x, err := Solve(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g\n", x)
	// Output:
	// [1 4 5]
}
