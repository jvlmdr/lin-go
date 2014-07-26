package lapack

import (
	"math"
	"sort"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestSVD(t *testing.T) {
	m, n := 150, 100
	want := randMat(m, n)

	// Take singular value decomposition.
	u, s, vt, err := SVD(want)
	if err != nil {
		t.Fatal(err)
	}

	// Check that A = U S V'.
	got := mat.Mul(u, mat.Mul(mat.NewDiag(s), vt))
	testMatEq(t, want, got)
}

func TestSVD_vsEig(t *testing.T) {
	m, n := 150, 100
	a := randMat(m, n)
	g := mat.Mul(mat.T(a), a)

	// Take eigen decomposition of Gram matrix.
	_, eigs, err := EigSymm(g)
	if err != nil {
		t.Fatal(err)
	}
	// Sort in descending order.
	sort.Sort(sort.Reverse(sort.Float64Slice(eigs)))
	// Take square root of eigenvalues.
	for i := range eigs {
		// Clip small negative values to zero.
		eigs[i] = math.Sqrt(math.Max(0, eigs[i]))
	}

	// Take singular value decomposition.
	_, svals, _, err := SVD(a)
	if err != nil {
		t.Fatal(err)
	}

	testSliceEq(t, eigs, svals)
}
