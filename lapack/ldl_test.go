package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"testing"
)

func TestLDLSolve(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(8, 4))
	A = mat.MakeCopy(mat.Times(A.T(), A))
	ldl, err := LDL(mat.MakeCopy(A))
	if err != nil {
		t.Fatal(err)
	}

	want := vec.MakeCopy(vec.Randn(4))
	got := vec.MakeCopy(mat.TimesVec(A, want))

	if err := ldl.Solve(mat.FromSlice(got)); err != nil {
		t.Fatal(err)
	}
	checkVectorsEqual(t, want, got, 1e-9)
}

func ExampleLDLSolve() {
	A := mat.MakeCopy(mat.Randn(2, 2))
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(1, 0, 2)
	A.Set(1, 1, -3)

	x := vec.Make(2)
	x.Set(0, 5)
	x.Set(1, -4)

	ldl, err := LDL(A)
	if err != nil {
		panic(err)
	}
	if err := ldl.Solve(mat.FromSlice(x)); err != nil {
		panic(err)
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
