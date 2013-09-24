package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
	"testing"
)

func TestLDLFact_Solve(t *testing.T) {
	// Random symmetric matrix.
	A := mat.MakeCopy(mat.Randn(4, 4))
	A = mat.MakeCopy(mat.Scale(0.5, mat.Plus(A, A.T())))
	ldl, err := LDL(A)
	if err != nil {
		t.Fatal(err)
	}

	// Random vector.
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))

	got, err := ldl.Solve(b)
	if err != nil {
		t.Fatal(err)
	}
	checkEqualVec(t, want, got, 1e-9)
}

func ExampleLDLFact_Solve() {
	A := mat.Make(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(1, 0, 2)
	A.Set(1, 1, -3)

	b := vec.Make(2)
	b.Set(0, 5)
	b.Set(1, -4)

	ldl, err := LDL(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := ldl.Solve(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
