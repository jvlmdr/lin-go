package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"testing"
)

func TestCholSolve(t *testing.T) {
	// Random symmetric positive definite matrix.
	A := mat.MakeCopy(mat.Randn(8, 4))
	A = mat.MakeCopy(mat.Times(A.T(), A))
	chol, err := Chol(A)
	if err != nil {
		t.Fatal(err)
	}

	// Random vector.
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))

	got, err := chol.Solve(b)
	if err != nil {
		t.Fatal(err)
	}
	checkEqualVec(t, want, got, 1e-9)
}

func ExampleCholFact_Solve() {
	A := mat.Make(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, 1)
	A.Set(1, 0, 1)
	A.Set(1, 1, 2)
	A = mat.MakeCopy(mat.Times(A.T(), A))

	var b vec.Slice = []float64{8, 13}

	chol, err := Chol(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := chol.Solve(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
