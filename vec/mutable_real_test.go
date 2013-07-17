package vec

import "fmt"

func ExampleCumSum() {
	x := Slice([]float64{1, 2, 3, 4})
	CumSum(x, x)
	fmt.Println(String(x))
	// Output:
	// (1, 3, 6, 10)
}
