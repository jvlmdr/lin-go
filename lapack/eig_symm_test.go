package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"sort"
	"testing"
)

func TestEigSymm(t *testing.T) {
	n := 100
	a := randMat(n, n)
	a = mat.Plus(a, mat.T(a))

	// Take eigen decomposition.
	v, d, err := EigSymm(a)
	if err != nil {
		t.Fatal(err)
	}

	got := mat.Mul(mat.Mul(v, mat.NewDiag(d)), mat.T(v))
	testMatEq(t, a, got)
}

func TestEigSymm_vsTrace(t *testing.T) {
	n := 100
	a := randMat(n, n)
	a = mat.Plus(a, mat.T(a))
	// Compute matrix trace.
	tr := mat.Tr(a)

	// Take eigen decomposition.
	_, d, err := EigSymm(a)
	if err != nil {
		t.Fatal(err)
	}
	// Sum eigenvalues.
	var sum float64
	for _, eig := range d {
		sum += eig
	}

	if !epsEq(tr, sum, eps) {
		t.Errorf("want %.4g, got %.4g", tr, sum)
	}
}

func ExampleEigSymm() {
	a := mat.NewRows([][]float64{
		{7, -2, 0},
		{-2, 6, -2},
		{0, -2, 5},
	})

	_, d, err := EigSymm(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Sort(sort.Float64Slice(d))
	fmt.Printf("%.6g\n", d)
	// Output:
	// [3 6 9]
}
