package vec

import "fmt"

func ExamplePlus() {
	x := Slice([]float64{1, 2, 3, 4})
	y := Slice([]float64{1, -1, 0, -7})
	Copy(x, Plus(x, y))
	fmt.Println(String(x))

	// Output:
	// (2, 1, 3, -3)
}
