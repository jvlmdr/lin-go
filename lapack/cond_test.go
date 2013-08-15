package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Find minimum-error solution to
//	x     = 3,
//	    y = 6,
//	x + y = 3.
func ExampleSolveCond_overdetermined() {
	A := mat.Make(3, 3)
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

	x, _, _, err := SolveCond(A, b, -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 4, 0)
}

// Find minimum-norm solution to
//	x     + z = 6,
//	    y + z = 9.
func ExampleSolveCond_underdetermined() {
	A := mat.Make(3, 3)
	b := vec.Make(3)

	A.Set(0, 0, 1)
	A.Set(0, 1, 0)
	A.Set(0, 2, 1)
	b.Set(0, 6)

	A.Set(1, 0, 0)
	A.Set(1, 1, 1)
	A.Set(1, 2, 1)
	b.Set(1, 9)

	x, _, _, err := SolveCond(A, b, -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vec.Sprintf("%.6g", x))
	// Output:
	// (1, 4, 5)
}
