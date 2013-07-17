package zvec

import (
	"math/cmplx"
	"testing"
)

func TestCumSum(t *testing.T) {
	const (
		n   = 80
		eps = 1e-9
	)

	var x Const = MakeCopy(Randn(n))
	var y Mutable = Make(n)
	CumSum(y, x)

	if cmplx.Abs(y.At(0)-x.At(0)) > eps {
		t.Errorf("First elements should be equal (want %g, got %g)", x.At(0), y.At(0))
	}
	sum := Sum(x)
	if cmplx.Abs(y.At(y.Size()-1)-sum) > eps {
		t.Errorf("Last element should equal sum (want %g, got %g)", sum, y.At(y.Size()-1))
	}
}
