package blas

import (
	"github.com/jackvalmadre/lin-go/zmat"
	"math/rand"
	"testing"
)

func TestComplexMatrixTimesMatrix(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeCopy(zmat.Randn(m, k))
	B := zmat.MakeCopy(zmat.Randn(n, k))

	got := ComplexMatrixTimesMatrix(alpha, A, NoTrans, B, Trans)
	want := zmat.Scale(alpha, zmat.Times(A, B.T()))
	zmat.CheckEqual(t, want, got, 1e-9)
}

func TestComplexMatrixTimesMatrixPlusMatrix(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeCopy(zmat.Randn(k, m))
	B := zmat.MakeCopy(zmat.Randn(k, n))
	C := zmat.MakeCopy(zmat.Randn(m, n))

	got := ComplexMatrixTimesMatrixPlusMatrix(alpha, A, Trans, B, NoTrans, C)
	want := zmat.Plus(zmat.Scale(alpha, zmat.Times(A.T(), B)), C)
	zmat.CheckEqual(t, want, got, 1e-9)
}

func TestComplexMatrixTimesMatrixPlusMatrixInPlace(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	beta := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeCopy(zmat.Randn(m, k))
	B := zmat.MakeCopy(zmat.Randn(k, n))
	C := zmat.MakeCopy(zmat.Randn(m, n))

	got := zmat.MakeCopy(C)
	ComplexMatrixTimesMatrixPlusMatrixInPlace(alpha, A, NoTrans, B, NoTrans, beta, got)
	want := zmat.Plus(zmat.Scale(alpha, zmat.Times(A, B)), zmat.Scale(beta, C))
	zmat.CheckEqual(t, want, got, 1e-9)
}

func ExampleComplexMatrixTimesMatrix() {
	A := zmat.Make(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, -1+1i)
	A.Set(1, 0, 2-1i)
	A.Set(1, 1, 3i)

	B := zmat.Make(2, 2)
	B.Set(0, 0, 2+2i)
	B.Set(0, 1, 3i)
	B.Set(1, 0, -1-1i)
	B.Set(1, 1, 2)

	C := ComplexMatrixTimesMatrix(1, A, NoTrans, B, ConjTrans)
	zmat.Printf(" % 3g", C)
	// Output:
	//  (  5 +1i) ( -3 +3i)
	//  ( 11 -6i) ( -1 +9i)
}
