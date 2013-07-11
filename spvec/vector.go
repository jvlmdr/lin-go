package spvec

// Superset of vec.Const.
type Const interface {
	// Dimension of vector.
	Size() int
	// Element accessor.
	At(i int) float64
	// Iterator over the elements.
	Elements() <-chan Element
}

type Element struct {
	Index int
	Value float64
}

// Superset of vec.Mutable.
type Mutable interface {
	Const
	// Overwrites or inserts an element.
	Set(i int, x float64)
	// Removes an element (different from setting to zero).
	Delete(i int)
}

// Copies a sparse vector only on its support.
func Copy(dst vec.Mutable, src Sparse) {
	ch := src.Elements()
	for elem := range ch {
		i, x := elem.Index, elem.Value
		dst.Set(i, dst.At(i), x)
	}
}
