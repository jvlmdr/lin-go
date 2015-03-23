package clap

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/cmat"
)

func TestCholFact_Solve(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = cmat.Mul(cmat.H(a), a)
	// Random vector.
	want := randVec(n)
	b := cmat.MulVec(a, want)

	// Factorize matrix.
	chol, err := Chol(a)
	if err != nil {
		t.Fatal(err)
	}

	got, err := chol.Solve(b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleCholFact_Solve() {
	// A = V' V, with V = [1, 1; 2, 1]
	v := cmat.NewRows([][]complex128{
		{1, 1},
		{2, 1},
	})
	a := cmat.Mul(cmat.H(v), v)

	// x = [1; 2]
	// b = V' V x
	//   = V' [1, 1; 2, 1] [1; 2]
	//   = [1, 2; 1, 1] [3; 4]
	//   = [11; 7]
	b := []complex128{11, 7}

	chol, err := Chol(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := chol.Solve(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(formatSlice(x, 'f', 3))
	// Output:
	// [(1.000+0.000i) (2.000+0.000i)]
}

func TestCholFact_Inverse(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = cmat.Mul(cmat.H(a), a)

	// Factorize matrix.
	chol, err := Chol(a)
	if err != nil {
		t.Fatal(err)
	}

	b, err := chol.Inv()
	if err != nil {
		t.Fatal(err)
	}
	right := cmat.Mul(a, b)
	testMatEq(t, cmat.I(n), right)
	left := cmat.Mul(a, b)
	testMatEq(t, cmat.I(n), left)
}
