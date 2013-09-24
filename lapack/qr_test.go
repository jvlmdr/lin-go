package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"fmt"
	"testing"
)

// Minimum-residual solution to over-constrained system by QR decomposition.
func TestQRFact_Solve_overdetermined(t *testing.T) {
	const (
		m = 8
		n = 4
	)

	// Random symmetric matrix.
	A := mat.MakeStrideCopy(mat.Randn(m, n))
	// Random vector.
	want := vec.MakeCopy(vec.Randn(n))
	// Product.
	b := vec.MakeCopy(mat.TimesVec(A, want))

	// Factorize.
	qr, err := QR(A)
	if err != nil {
		t.Fatal(err)
	}
	// Solve.
	got, err := qr.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}

	checkEqualVec(t, want, got, 1e-9)
}

// Minimum-norm solution to under-constrained system by QR decomposition.
func TestQRFact_Solve_underdetermined(t *testing.T) {
	const (
		m = 8
		n = 4
	)

	// Random symmetric matrix.
	A := mat.MakeStrideCopy(mat.Randn(m, n))
	// Random vector.
	in := vec.MakeCopy(vec.Randn(m))
	// Product.
	b := vec.MakeCopy(mat.TimesVec(A.T(), in))

	// Factorize.
	qr, err := QR(A)
	if err != nil {
		t.Fatal(err)
	}
	// Solve.
	got, err := qr.Solve(true, b)
	if err != nil {
		t.Fatal(err)
	}

	// Check outputs are equal.
	checkEqualVec(t, b, mat.TimesVec(A.T(), got), 1e-9)
	// Check norm is not larger.
	if vec.Norm(got) > vec.Norm(in) {
		t.Fatalf("norm of solution was larger (want %g <= %g)", vec.Norm(got), vec.Norm(in))
	}
}

// [4, 2; 1, 1; 2, 0] \ [6; 7; 4] = [1; 2]
func ExampleQRFact_Solve_overdetermined() {
	A := mat.MakeStride(3, 2)
	A.Set(0, 0, 4)
	A.Set(0, 1, 2)
	A.Set(1, 0, 1)
	A.Set(1, 1, 1)
	A.Set(2, 0, 2)
	A.Set(2, 1, 0)
	b := vec.Slice([]float64{6, 7, 4})

	qr, err := QR(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := qr.Solve(false, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 2)
}

// pinv([4, 1, 2; 2, 1, 0]) * [39; 19] = [8; 3; 2]
func ExampleQRFact_Solve_underdetermined() {
	A := mat.MakeStride(3, 2)
	A.Set(0, 0, 4)
	A.Set(0, 1, 2)
	A.Set(1, 0, 1)
	A.Set(1, 1, 1)
	A.Set(2, 0, 2)
	A.Set(2, 1, 0)
	b := vec.Slice([]float64{39, 19})

	qr, err := QR(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := qr.Solve(true, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (8, 3, 2)
}

////////////////////////////////////////////////////////////////////////////////

// Minimum-residual solution to over-constrained system by QR decomposition.
func TestQRFactCmplx_Solve_overdetermined(t *testing.T) {
	const (
		m = 8
		n = 4
	)

	// Random symmetric matrix.
	A := zmat.MakeStrideCopy(zmat.Randn(m, n))
	// Random vector.
	want := zvec.MakeCopy(zvec.Randn(n))
	// Product.
	b := zvec.MakeCopy(zmat.TimesVec(A, want))

	// Factorize.
	qr, err := QRCmplx(A)
	if err != nil {
		t.Fatal(err)
	}
	// Solve.
	got, err := qr.Solve(false, b)
	if err != nil {
		t.Fatal(err)
	}

	checkEqualVecCmplx(t, want, got, 1e-9)
}

// [4, 2; 1, 1; 2, 0] \ [6; 7; 4] = [1; 2]
func ExampleQRFactCmplx_Solve_overdeterminedReal() {
	A := zmat.MakeStride(3, 2)
	A.Set(0, 0, 4)
	A.Set(0, 1, 2)
	A.Set(1, 0, 1)
	A.Set(1, 1, 1)
	A.Set(2, 0, 2)
	A.Set(2, 1, 0)
	b := zvec.Slice([]complex128{6, 7, 4})

	qr, err := QRCmplx(A)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := qr.Solve(false, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(zvec.Sprintf("%.6g", x))
	// Output:
	// ((1+0i), (2+0i))
}
