package vec

// Augment a slice with the required methods.
type Slice []float64

// Constructs a simple slice vector.
func MakeSlice(n int) Slice { return make([]float64, n) }

func (s Slice) Size() int            { return len(s) }
func (s Slice) At(i int) float64     { return s[i] }
func (s Slice) Set(i int, v float64) { s[i] = v }

// Constructs a simple slice vector and copies x into it.
func MakeSliceCopy(x Const) Slice {
	y := MakeSlice(x.Size())
	Copy(y, x)
	return y
}
