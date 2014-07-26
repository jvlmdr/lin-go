package blas

import (
	"math/rand"
	"testing"

	"github.com/jvlmdr/lin-go/zmat"
)

func TestMatTimesMatCmplx(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeStrideCopy(zmat.Randn(m, k))
	B := zmat.MakeStrideCopy(zmat.Randn(n, k))

	got := MatTimesMatCmplx(alpha, A, NoTrans, B, Trans)
	want := zmat.Scale(alpha, zmat.Times(A, B.T()))
	checkEqualMatCmplx(t, want, got, 1e-9)
}

func TestMatTimesMatPlusMatCmplx(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeStrideCopy(zmat.Randn(k, m))
	B := zmat.MakeStrideCopy(zmat.Randn(k, n))
	C := zmat.MakeStrideCopy(zmat.Randn(m, n))

	got := MatTimesMatPlusMatCmplx(alpha, A, Trans, B, NoTrans, C)
	want := zmat.Plus(zmat.Scale(alpha, zmat.Times(A.T(), B)), C)
	checkEqualMatCmplx(t, want, got, 1e-9)
}

func TestMatTimesMatPlusMatCmplxNoCopy(t *testing.T) {
	const (
		m = 3
		n = 5
		k = 4
	)

	alpha := complex(rand.NormFloat64(), rand.NormFloat64())
	beta := complex(rand.NormFloat64(), rand.NormFloat64())
	A := zmat.MakeStrideCopy(zmat.Randn(m, k))
	B := zmat.MakeStrideCopy(zmat.Randn(k, n))
	C := zmat.MakeStrideCopy(zmat.Randn(m, n))

	got := zmat.MakeStrideCopy(C)
	MatTimesMatPlusMatCmplxNoCopy(alpha, A, NoTrans, B, NoTrans, beta, got)
	want := zmat.Plus(zmat.Scale(alpha, zmat.Times(A, B)), zmat.Scale(beta, C))
	checkEqualMatCmplx(t, want, got, 1e-9)
}

func ExampleMatTimesMatCmplx() {
	A := zmat.MakeStride(2, 2)
	A.Set(0, 0, 1)
	A.Set(0, 1, -1+1i)
	A.Set(1, 0, 2-1i)
	A.Set(1, 1, 3i)

	B := zmat.MakeStride(2, 2)
	B.Set(0, 0, 2+2i)
	B.Set(0, 1, 3i)
	B.Set(1, 0, -1-1i)
	B.Set(1, 1, 2)

	C := MatTimesMatCmplx(1, A, NoTrans, B, ConjTrans)
	zmat.Printf(" % 3g", C)
	// Output:
	//  (  5 +1i) ( -3 +3i)
	//  ( 11 -6i) ( -1 +9i)
}
