package lapack

//	import (
//		"github.com/jackvalmadre/lin-go/mat"
//		"github.com/jackvalmadre/lin-go/vec"
//
//		"fmt"
//		"math"
//		"sort"
//		"testing"
//	)
//
//	// eig([7, -2, 0; -2, 6, -2; 0, -2, 5])
//	func ExampleEigSymm() {
//		A := mat.MakeStride(3, 3)
//		A.Set(0, 0, 7)
//		A.Set(0, 1, -2)
//		A.Set(0, 2, 0)
//		A.Set(1, 0, -2)
//		A.Set(1, 1, 6)
//		A.Set(1, 2, -2)
//		A.Set(2, 0, 0)
//		A.Set(2, 1, -2)
//		A.Set(2, 2, 5)
//
//		v, err := EigSymm(A)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//
//		sort.Sort(sort.Float64Slice(v))
//		fmt.Printf("(%.6g, %.6g, %.6g)\n", v[0], v[1], v[2])
//		// Output:
//		// (3, 6, 9)
//	}
//
//	func TestEigSymm_SumVsTrace(t *testing.T) {
//		const n = 100
//		A := mat.MakeStrideCopy(mat.Randn(n, n))
//
//		DA := mat.MakeStride(n, n)
//		d := vec.MakeSliceCopy(vec.Randn(n))
//		for i := 0; i < n; i++ {
//			d.Set(i, d.At(i)*float64(n-i)/float64(n))
//		}
//		for i := 0; i < n; i++ {
//			for j := 0; j < n; j++ {
//				DA.Set(i, j, d.At(i)*A.At(i, j))
//			}
//		}
//		A = mat.MakeStrideCopy(mat.Times(A.T(), DA))
//
//		eigs, err := EigSymm(A)
//		if err != nil {
//			t.Fatal(err)
//		}
//
//		var tr float64
//		for i := 0; i < n; i++ {
//			tr += A.At(i, i)
//		}
//
//		var sum float64
//		for _, eig := range eigs {
//			sum += eig
//		}
//
//		const eps = 0
//		if math.Abs(sum-tr) > eps {
//			t.Errorf("trace %g, sum eigs %g", tr, sum)
//		}
//	}
