package blas

import (
	"math/rand"
	"testing"

	"github.com/jvlmdr/lin-go/mat"
)

func TestMatTimesMat(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	A := mat.MakeStrideCopy(mat.Randn(m, k))
	B := mat.MakeStrideCopy(mat.Randn(n, k))

	got := MatTimesMat(alpha, A, NoTrans, B, Trans)
	want := mat.Scale(alpha, mat.Times(A, B.T()))
	checkEqualMat(t, want, got, 1e-9)
}

func TestMatTimesMatPlusMat(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	A := mat.MakeStrideCopy(mat.Randn(k, m))
	B := mat.MakeStrideCopy(mat.Randn(k, n))
	C := mat.MakeStrideCopy(mat.Randn(m, n))

	got := MatTimesMatPlusMat(alpha, A, Trans, B, NoTrans, C)
	want := mat.Plus(mat.Scale(alpha, mat.Times(A.T(), B)), C)
	checkEqualMat(t, want, got, 1e-9)
}

func TestMatTimesMatPlusMatNoCopy(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	beta := rand.NormFloat64()
	A := mat.MakeStrideCopy(mat.Randn(m, k))
	B := mat.MakeStrideCopy(mat.Randn(k, n))
	C := mat.MakeStrideCopy(mat.Randn(m, n))

	got := mat.MakeStrideCopy(C)
	MatTimesMatPlusMatNoCopy(alpha, A, NoTrans, B, NoTrans, beta, got)
	want := mat.Plus(mat.Scale(alpha, mat.Times(A, B)), mat.Scale(beta, C))
	checkEqualMat(t, want, got, 1e-9)
}

func ExampleMatTimesMat() {
	A := mat.MakeStride(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, -1)
	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 2)

	B := mat.MakeStride(3, 2)
	B.Set(0, 0, 1)
	B.Set(0, 1, 0)
	B.Set(1, 0, -1)
	B.Set(1, 1, 0)
	B.Set(2, 0, 1)
	B.Set(2, 1, 2)

	C := MatTimesMat(1, A, NoTrans, B, NoTrans)
	mat.Printf("% 3g", C)
	// Output:
	//   0 -2
	//   1  4
}

func ExampleMatTimesMatPlusMatNoCopy() {
	A := mat.MakeStride(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, -1)
	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 2)

	B := mat.MakeStride(3, 2)
	B.Set(0, 0, 1)
	B.Set(0, 1, 0)
	B.Set(1, 0, -1)
	B.Set(1, 1, 0)
	B.Set(2, 0, 1)
	B.Set(2, 1, 2)

	C := mat.MakeStride(2, 2)
	C.Set(0, 0, -4)
	C.Set(0, 1, -3)
	C.Set(1, 0, 2)
	C.Set(1, 1, 1)

	MatTimesMatPlusMatNoCopy(1, A, NoTrans, B, NoTrans, -1, C)
	mat.Printf("% 3g", C)
	// Output:
	//   4  1
	//  -1  3
}
