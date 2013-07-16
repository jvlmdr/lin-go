package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"testing"
)

// Find minimum-error solution to
//	x     = 3,
//	    y = 6,
//	x + y = 3.
func ExampleFullRankSolve_overdetermined() {
	A := mat.Make(3, 2)
	b := vec.Make(3)

	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	b.Set(0, 3)

	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	b.Set(1, 6)

	A.Set(2, 0, 1)
	A.Set(2, 1, 1)
	b.Set(2, 3)

	x := FullRankSolve(A, b)
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 4)
}

// Find minimum-norm solution to
//	x     + z = 6,
//	    y + z = 9.
func ExampleFullRankSolve_underdetermined() {
	A := mat.Make(2, 3)
	b := vec.Make(2)

	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, 1)
	b.Set(0, 6)

	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 1)
	b.Set(1, 9)

	x := FullRankSolve(A, b)
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 4, 5)
}

// Compare against pseudo-inverse solution.
func TestFullRankOverdetermined(t *testing.T) {
	const (
		m = 8
		n = 5
		p = 3
	)
	A := mat.MakeCopy(mat.Randn(m, n))
	B := mat.MakeCopy(mat.Randn(m, p))
	got := FullRankSolveMatrix(A, B)

	AA := mat.MakeCopy(mat.Times(mat.T(A), A))
	AB := mat.MakeCopy(mat.Times(mat.T(A), B))
	want, _ := SquareSolveMatrix(AA, AB)

	mat.CheckEqual(t, want, got, 1e-9)
}

// Compare against pseudo-inverse solution.
func TestFullRankUnderdetermined(t *testing.T) {
	const (
		m = 5
		n = 8
		p = 3
	)
	A := mat.MakeCopy(mat.Randn(m, n))
	B := mat.MakeCopy(mat.Randn(m, p))
	got := FullRankSolveMatrix(A, B)

	AA := mat.MakeCopy(mat.Times(A, mat.T(A)))
	Y, _ := SquareSolveMatrix(AA, B)
	want := mat.MakeCopy(mat.Times(mat.T(A), Y))

	mat.CheckEqual(t, want, got, 1e-9)
}
