package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
	"testing"
)

func TestLUFact_Solve(t *testing.T) {
	// Random symmetric matrix.
	A := mat.MakeStrideCopy(mat.Randn(4, 4))
	A = mat.MakeStrideCopy(mat.Scale(0.5, mat.Plus(A, A.T())))
	lu, err := LU(A)
	if err != nil {
		t.Fatal(err)
	}

	// Random vector.
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))

	got, err := lu.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}
	checkVectorsEqual(t, want, got, 1e-9)
}

func ExampleLUFact_Solve() {
	A := mat.MakeStride(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(1, 0, 2)
	A.Set(1, 1, -3)

	b := vec.Make(2)
	b.Set(0, 5)
	b.Set(1, -4)

	lu, err := LU(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := lu.Solve(false, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
