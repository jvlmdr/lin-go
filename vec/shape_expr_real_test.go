package vec

import "fmt"

func ExampleCat() {
	var (
		a = Slice([]float64{1, 2})
		b = Slice([]float64{3})
		c = Slice([]float64{})
		d = Slice([]float64{4, 5, 6})
	)
	fmt.Println(String(Cat(a, b, c, d)))
	// Output:
	// (1, 2, 3, 4, 5, 6)
}

func ExampleMutableCat() {
	var (
		a = MakeSlice(4)
		b = MakeSlice(3)
		c = Slice([]float64{1, 2, 3, 4, 5, 6, 7})
	)
	Copy(MutableCat(a, b), c)
	fmt.Println(String(a))
	fmt.Println(String(b))
	// Output:
	// (1, 2, 3, 4)
	// (5, 6, 7)
}
