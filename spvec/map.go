package spvec

type MapVector struct {
	N   int
	Map map[int]float64
}

func (v MapVector) Size() int {
	return v.N
}

func (v MapVector) At(i int) float64 {
	return v.Map[i]
}

func (v MapVector) Elements() <-chan Element {
	ch := make(chan Element)
	go func() {
		for i, x := range v.Map {
			ch <- Element{i, x}
		}
		close(ch)
	}()
	return ch
}

func (v MapVector) Set(i int, x float64) {
	v.Map[i] = x
}

func (v MapVector) Delete(i int) {
	delete(v.Map, i)
}
