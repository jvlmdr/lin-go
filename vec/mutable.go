package vec

// Miscellaneous operations you can do with a Mutable vector.

import "fmt"

func Copy(dst Mutable, src Const) {
	if dst.Size() != src.Size() {
		panic(fmt.Sprintf("Vectors are different sizes (%d and %d)", dst.Size(), src.Size()))
	}
	for i := 0; i < src.Size(); i++ {
		dst.Set(i, src.At(i))
	}
}
