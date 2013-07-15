package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"math"
	"math/cmplx"
	"testing"
)

func ExampleSquareSolve() {
	A := mat.MakeCopy(mat.Randn(4, 4))
	z := vec.Make(4)
	z.Set(0, 1)
	z.Set(1, 2)
	z.Set(2, 3)
	z.Set(3, 4)

	b := vec.MakeCopy(mat.TimesVec(A, z))
	x, _ := SquareSolve(A, b)
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2, 3, 4)
}

func ExampleSquareSolveInPlace() {
	A := mat.MakeCopy(mat.Randn(4, 4))
	z := vec.Make(4)
	z.Set(0, 1)
	z.Set(1, 2)
	z.Set(2, 3)
	z.Set(3, 4)

	x := vec.MakeSliceCopy(mat.TimesVec(A, z))
	SquareSolveInPlace(A, x)
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2, 3, 4)
}

func TestSquareSolve(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(4, 4))
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))
	got, _ := SquareSolve(A, b)
	checkVectorsEqual(t, want, got, 1e-6)
}

func TestSquareSolveInPlace(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(4, 4))
	want := vec.MakeCopy(vec.Randn(4))
	got := vec.MakeSliceCopy(mat.TimesVec(A, want))
	SquareSolveInPlace(A, got)
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

func ExampleSquareSolveComplex() {
	A := zmat.MakeCopy(zmat.Randn(4, 4))
	z := zvec.Make(4)
	z.Set(0, 1+1i)
	z.Set(1, 2+2i)
	z.Set(2, 3+3i)
	z.Set(3, 4+4i)

	b := zvec.MakeCopy(zmat.TimesVec(A, z))
	x, _ := SquareSolveComplex(A, b)
	fmt.Println(zvec.Sprintf("%.6g", x))
	// Output:
	// ((1+1i), (2+2i), (3+3i), (4+4i))
}

func checkComplexVectorsEqual(t *testing.T, want, got zvec.Const, eps float64) {
	if want.Size() != got.Size() {
		t.Fatalf("Vector sizes do not match (want %d, got %d)", want.Size(), got.Size())
	}

	for i := 0; i < want.Size(); i++ {
		x := want.At(i)
		y := got.At(i)
		if cmplx.Abs(x-y) > 1e-6 {
			t.Errorf("Wrong value at index %d (want %g, got %g)", i, x, y)
		}
	}
}
