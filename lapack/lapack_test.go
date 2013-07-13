package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/mat"
	"testing"
)

func TestSolveVec(t *testing.T) {
	A := mat.MakeCopy(mat.Randn(4, 4))
	want := vec.MakeCopy(vec.Randn(4))
	b := vec.MakeCopy(mat.TimesVec(A, want))
	got := Solve(A, b)

	for i := 0; i < want.Size(); i++ {
		fmt.Println(want.At(i), got.At(i))
	}
}
