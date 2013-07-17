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

func ExampleMultiply() {
	x := Slice([]float64{1, 2, 3, 4})
	y := Slice([]float64{1, -1, 0, -7})
	z := Slice([]float64{1, 3, 3, -9})
	Copy(x, Multiply(x, y, z))
	fmt.Println(String(x))

	// Output:
	// (1, -6, 0, 252)
}
