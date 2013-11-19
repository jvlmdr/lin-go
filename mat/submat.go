package mat

type Elem struct {
	I, J int
}

func (a Elem) Add(b Elem) Elem {
	return Elem{a.I + b.I, a.J + b.J}
}

func (a Elem) Sub(b Elem) Elem {
	return Elem{a.I - b.I, a.J - b.J}
}

func (e Elem) In(r Rectangle) bool {
	return r.Min.I <= e.I && e.I < r.Max.I && r.Min.J <= e.J && e.J < r.Max.J
}

type Rectangle struct {
	Min, Max Elem
}

func (r Rectangle) Dims() (int, int) {
	s := r.Max.Sub(r.Min)
	return s.I, s.J
}

func (r Rectangle) In(q Rectangle) bool {
	return r.Min.In(q) && r.Max.Sub(Elem{1, 1}).In(q)
}

// Copies a submatrix from a matrix.
func Submat(a Const, r Rectangle) Mutable {
	// Check bounds.
	p, q := a.Dims()
	bnds := Rectangle{Elem{0, 0}, Elem{p, q}}
	if !r.In(bnds) {
		panic(errRectOutsideMat(r, p, q))
	}

	m, n := r.Dims()
	b := New(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			e := Elem{i, j}.Add(r.Min)
			b.Set(i, j, a.At(e.I, e.J))
		}
	}
	return b
}

// Mutable reference to a submatrix within a matrix.
type SubmatRef struct {
	Mat Mutable
	Rectangle
}

func (a SubmatRef) At(i, j int) float64 {
	e := Elem{i, j}.Add(a.Min)
	return a.At(e.I, e.J)
}

func (a SubmatRef) Set(i, j int, v float64) {
	e := Elem{i, j}.Add(a.Min)
	a.Set(e.I, e.J, v)
}
