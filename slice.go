package vec

type Slice []float64

func NewSlice(n int) Slice           { return make([]float64, n) }
func (s Slice) Size() int            { return len(s) }
func (s Slice) At(i int) float64     { return s[i] }
func (s Slice) Set(i int, v float64) { s[i] = v }

func CopyToNewSlice(x Const) Slice {
	y := NewSlice(x.Size())
	CopyTo(y, x)
	return y
}
