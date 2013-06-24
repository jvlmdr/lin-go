package vec

type Slice []float64

func (s Slice) Size() int {
	return len(s)
}

func (s Slice) At(i int) float64 {
	return s[i]
}

func (s Slice) Set(i int, v float64) {
	s[i] = v
}
