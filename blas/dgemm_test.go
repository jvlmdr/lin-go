package blas

import (
	"github.com/jackvalmadre/lin-go/mat"
	"math/rand"
	"testing"
)

func TestMatrixTimesMatrix(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	A := mat.MakeCopy(mat.Randn(m, k))
	B := mat.MakeCopy(mat.Randn(n, k))

	got := MatrixTimesMatrix(alpha, A, NoTrans, B, Trans)
	want := mat.Scale(alpha, mat.Times(A, B.T()))
	mat.CheckEqual(t, want, got, 1e-9)
}

func TestMatrixTimesMatrixPlusMatrix(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	A := mat.MakeCopy(mat.Randn(k, m))
	B := mat.MakeCopy(mat.Randn(k, n))
	C := mat.MakeCopy(mat.Randn(m, n))

	got := MatrixTimesMatrixPlusMatrix(alpha, A, Trans, B, NoTrans, C)
	want := mat.Plus(mat.Scale(alpha, mat.Times(A.T(), B)), C)
	mat.CheckEqual(t, want, got, 1e-9)
}

func TestMatrixTimesMatrixPlusMatrixInPlace(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := rand.NormFloat64()
	beta := rand.NormFloat64()
	A := mat.MakeCopy(mat.Randn(m, k))
	B := mat.MakeCopy(mat.Randn(k, n))
	C := mat.MakeCopy(mat.Randn(m, n))

	got := mat.MakeCopy(C)
	MatrixTimesMatrixPlusMatrixInPlace(alpha, A, NoTrans, B, NoTrans, beta, got)
	want := mat.Plus(mat.Scale(alpha, mat.Times(A, B)), mat.Scale(beta, C))
	mat.CheckEqual(t, want, got, 1e-9)
}

func ExampleMatrixTimesMatrix() {
	A := mat.Make(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, -1)
	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 2)

	B := mat.Make(3, 2)
	B.Set(0, 0, 1)
	B.Set(0, 1, 0)
	B.Set(1, 0, -1)
	B.Set(1, 1, 0)
	B.Set(2, 0, 1)
	B.Set(2, 1, 2)

	C := MatrixTimesMatrix(1, A, NoTrans, B, NoTrans)
	mat.Printf("% 3g", C)
	// Output:
	//   0 -2
	//   1  4
}

func ExampleMatrixTimesMatrixPlusMatrixInPlace() {
	A := mat.Make(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, -1)
	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 2)

	B := mat.Make(3, 2)
	B.Set(0, 0, 1)
	B.Set(0, 1, 0)
	B.Set(1, 0, -1)
	B.Set(1, 1, 0)
	B.Set(2, 0, 1)
	B.Set(2, 1, 2)

	C := mat.Make(2, 2)
	C.Set(0, 0, -4)
	C.Set(0, 1, -3)
	C.Set(1, 0, 2)
	C.Set(1, 1, 1)

	MatrixTimesMatrixPlusMatrixInPlace(1, A, NoTrans, B, NoTrans, -1, C)
	mat.Printf("% 3g", C)
	// Output:
	//   4  1
	//  -1  3
}
