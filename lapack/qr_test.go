package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

// Minimum-residual solution to over-constrained system by QR decomposition.
func TestQRFact_Solve_overdetermined(t *testing.T) {
	m, n := 150, 100
	a, b, want, err := overDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	qr, err := QR(a)
	if err != nil {
		t.Fatal(err)
	}
	got, err := qr.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

// Minimum-norm solution to under-constrained system by QR decomposition.
func TestQRFact_Solve_underdetermined(t *testing.T) {
	m, n := 100, 150
	a, b, want, err := underDetProb(m, n)
	if err != nil {
		t.Fatal(err)
	}

	// Take QR factorization of transpose.
	qr, err := QR(mat.T(a))
	if err != nil {
		t.Fatal(err)
	}
	got, err := qr.Solve(true, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleQRFact_Solve() {
	a := mat.NewRows([][]float64{
		{4, 2},
		{1, 1},
		{2, 0},
	})
	b_over := []float64{6, 7, 4}
	b_under := []float64{39, 19}

	qr, err := QR(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	x_over, err := qr.Solve(false, b_over)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g\n", x_over)

	x_under, err := qr.Solve(true, b_under)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g\n", x_under)
	// Output:
	// [1 2]
	// [8 3 2]
}
