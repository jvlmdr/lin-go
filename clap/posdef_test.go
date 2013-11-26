package clap

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/cmat"
	"testing"
)

func TestSolvePosDef(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = cmat.Mul(cmat.H(a), a)
	// Random vector.
	want := randVec(n)
	b := cmat.MulVec(a, want)

	got, err := SolvePosDef(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolvePosDef() {
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

	x, err := SolvePosDef(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [(1+0i) (2+0i)]
}
