package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"testing"
)

func TestLDLSolve(t *testing.T) {
	// Random symmetric matrix.
	A := mat.MakeCopy(mat.Randn(4, 4))
	A = mat.MakeCopy(mat.Scale(0.5, mat.Plus(A, A.T())))
	ldl, err := LDL(mat.MakeCopy(A))
	if err != nil {
		t.Fatal(err)
	}

	// Random vector.
	want := vec.MakeCopy(vec.Randn(4))
	got := vec.MakeCopy(mat.TimesVec(A, want))

	if err := ldl.Solve(mat.FromSlice(got)); err != nil {
		t.Fatal(err)
	}
	checkVectorsEqual(t, want, got, 1e-9)
}

func ExampleLDLSolve() {
	A := mat.Make(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(1, 0, 2)
	A.Set(1, 1, -3)

	x := vec.Make(2)
	x.Set(0, 5)
	x.Set(1, -4)

	ldl, err := LDL(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := ldl.Solve(mat.FromSlice(x)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
