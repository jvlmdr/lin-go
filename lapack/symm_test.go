package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"testing"
)

func TestSolveSymm(t *testing.T) {
	// Random symmetric matrix.
	A := mat.MakeCopy(mat.Randn(4, 4))
	A = mat.MakeCopy(mat.Scale(0.5, mat.Plus(A, A.T())))
	// Random vector.
	x := vec.MakeCopy(vec.Randn(4))

	b := vec.MakeCopy(mat.TimesVec(A, x))
	got, err := SolveSymm(A, b)
	if err != nil {
		t.Fatal(err)
	}
	checkEqualVec(t, x, got, 1e-9)
}

func ExampleSolveSymm() {
	A := mat.Make(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(1, 0, 2)
	A.Set(1, 1, -3)

	b := vec.Make(2)
	b.Set(0, 5)
	b.Set(1, -4)

	x, err := SolveSymm(A, b)
	if err != nil {
		panic(err)
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}
