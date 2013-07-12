package mat

type Size struct{ Rows, Cols int }

func (s Size) Equals(t Size) bool {
	return s.Rows == t.Rows && s.Cols == t.Cols
}

// Returns Rows == Cols.
func (s Size) Square() bool {
	return s.Rows == s.Cols
}

// Returns the transposed size.
func (s Size) T() Size {
	return Size{s.Cols, s.Rows}
}

// Returns Rows * Cols.
func (s Size) Area() int {
	return s.Rows * s.Cols
}

type Index struct{ I, J int }

// Returns the transposed index.
func (p Index) T() Index {
	return Index{p.J, p.I}
}

type Rect struct{ Min, Max Index }

func MakeRect(i0, j0, i1, j1 int) Rect {
	return Rect{Index{i0, j0}, Index{i1, j1}}
}

func (r Rect) Size() Size {
	return Size{r.Rows(), r.Cols()}
}

func (r Rect) Rows() int {
	return r.Max.I - r.Min.I
}

func (r Rect) Cols() int {
	return r.Max.J - r.Min.J
}

func (r Rect) T() Rect {
	return Rect{r.Min.T(), r.Max.T()}
}
