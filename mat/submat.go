package mat

// Copies the contents of src to dst.
//
// Panics if the two matrices are not the same size.
func Copy(dst Mutable, src Const) {
	if err := errIfDimsNotEq(src, dst); err != nil {
		panic(err)
	}

	m, n := src.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dst.Set(i, j, src.At(i, j))
		}
	}
}

type Elem struct {
	I, J int
}

func (a Elem) Add(b Elem) Elem {
	return Elem{a.I + b.I, a.J + b.J}
}

func (a Elem) Sub(b Elem) Elem {
	return Elem{a.I - b.I, a.J - b.J}
}

// Tests if an element (i, j) is inside the rectangle.
func (e Elem) In(r Rectangle) bool {
	return r.Min.I <= e.I && e.I < r.Max.I && r.Min.J <= e.J && e.J < r.Max.J
}

// All points (i, j) such that
//	Min.I <= i < Max.I
//	Min.J <= j < Max.J
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

// Creates a copy of a submatrix within a matrix.
func Sub(a Const, r Rectangle) *Mat {
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
// Idiomatic use:
//	r := mat.Rect(imin, jmin, imax, jmax)
//	mat.Copy(mat.Ref{x, r}, y)
type Ref struct {
	Mat Mutable
	Rectangle
}

func (a Ref) At(i, j int) float64 {
	e := Elem{i, j}.Add(a.Min)
	return a.Mat.At(e.I, e.J)
}

func (a Ref) Set(i, j int, v float64) {
	e := Elem{i, j}.Add(a.Min)
	a.Mat.Set(e.I, e.J, v)
}
