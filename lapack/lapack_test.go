package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"math"
	"testing"
)

func ExampleSolve() {
	A := mat.MakeCopy(mat.Randn(4, 4))
	z := vec.Make(4)
	z.Set(0, 0)
	z.Set(1, 1)
	z.Set(2, 2)
	z.Set(3, 3)

	b := vec.MakeCopy(mat.TimesVec(A, z))
	x := Solve(A, b)
	fmt.Println(vec.Sprintf("%.3f", x))
	// Output:
	// (0.000, 1.000, 2.000, 3.000)
}

func ExampleSolveInPlace() {
	A := mat.MakeCopy(mat.Randn(4, 4))
	z := vec.Make(4)
	z.Set(0, 0)
	z.Set(1, 1)
	z.Set(2, 2)
	z.Set(3, 3)

	x := vec.MakeSliceCopy(mat.TimesVec(A, z))
	SolveInPlace(A, x)
	fmt.Println(vec.Sprintf("%.3f", x))
	// Output:
	// (0.000, 1.000, 2.000, 3.000)
}

func TestSolve(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(4, 4))
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))
	got := Solve(A, b)
	checkVectorsEqual(t, want, got, 1e-6)
}

func TestSolveInPlace(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(4, 4))
	want := vec.MakeCopy(vec.Randn(4))
	got := vec.MakeSliceCopy(mat.TimesVec(A, want))
	SolveInPlace(A, got)
	checkVectorsEqual(t, want, got, 1e-6)
}

func checkVectorsEqual(t *testing.T, want, got vec.Const, eps float64) {
	if want.Size() != got.Size() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Size(), got.Size())
	}

	for i := 0; i < want.Size(); i++ {
		x := want.At(i)
		y := got.At(i)
		if math.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}
