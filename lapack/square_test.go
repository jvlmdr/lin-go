package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
	"testing"
)

func ExampleSolveSquare() {
	A := mat.MakeStrideCopy(mat.Randn(4, 4))
	z := vec.MakeSlice(4)
	z.Set(0, 1)
	z.Set(1, 2)
	z.Set(2, 3)
	z.Set(3, 4)

	b := vec.MakeCopy(mat.TimesVec(A, z))
	x, _ := SolveSquare(A, b)
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2, 3, 4)
}

func TestSolveSquare(t *testing.T) {
	const n = 100
	A := mat.MakeStrideCopy(mat.Randn(n, n))
	want := vec.MakeCopy(vec.Randn(n))
	b := vec.MakeCopy(mat.TimesVec(A, want))
	got, _ := SolveSquare(A, b)
	checkEqualVec(t, want, got, 1e-9)
}

func ExampleSolveSquareCmplx() {
	A := zmat.MakeStrideCopy(zmat.Randn(4, 4))
	z := zvec.Make(4)
	z.Set(0, 1+1i)
	z.Set(1, 2+2i)
	z.Set(2, 3+3i)
	z.Set(3, 4+4i)

	b := zvec.MakeCopy(zmat.TimesVec(A, z))
	x, _ := SolveSquareCmplx(A, b)
	fmt.Println(zvec.Sprintf("%.6g", x))
	// Output:
	// ((1+1i), (2+2i), (3+3i), (4+4i))
}
