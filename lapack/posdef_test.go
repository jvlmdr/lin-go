package lapack

import (
	"fmt"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestSolvePosDef(t *testing.T) {
	n := 100
	// Random symmetric positive definite matrix.
	a := randMat(2*n, n)
	a = mat.Mul(mat.T(a), a)
	// Random vector.
	want := randVec(n)
	b := mat.MulVec(a, want)

	got, err := SolvePosDef(a, b)
	if err != nil {
		t.Fatal(err)
	}
	testSliceEq(t, want, got)
}

func ExampleSolvePosDef() {
	// A = V' V, with V = [1, 1; 2, 1]
	v := mat.NewRows([][]float64{
		{1, 1},
		{2, 1},
	})
	a := mat.Mul(mat.T(v), v)

	// x = [1; 2]
	// b = V' V x
	//   = V' [1, 1; 2, 1] [1; 2]
	//   = [1, 2; 1, 1] [3; 4]
	//   = [11; 7]
	b := []float64{11, 7}

	x, err := SolvePosDef(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6g", x)
	// Output:
	// [1 2]
}
