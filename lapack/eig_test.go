package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"

	"fmt"
	"sort"
)

// eig([7, -2, 0; -2, 6, -2; 0, -2, 5])
func ExampleEigSymm() {
	A := mat.MakeStride(3, 3)
	A.Set(0, 0, 7)
	A.Set(0, 1, -2)
	A.Set(0, 2, 0)
	A.Set(1, 0, -2)
	A.Set(1, 1, 6)
	A.Set(1, 2, -2)
	A.Set(2, 0, 0)
	A.Set(2, 1, -2)
	A.Set(2, 2, 5)

	v, err := EigSymm(A)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Sort(sort.Float64Slice(v))
	fmt.Printf("(%.6g, %.6g, %.6g)\n", v[0], v[1], v[2])
	// Output:
	// (3, 6, 9)
}
