package blas

// When RowMaj is false, Stride is the column stride.
// It is possible for the column stride to be greater than the number of rows.
// If RowMaj is true then the elements are effectively stored in row-major order,
// but with the guarantee that the row stride is 1.
type Mat struct {
	Rows   int
	Cols   int
	Stride int
	RowMaj bool
	// If RowMaj is false, then element (i, j) is at Elems[i + j*Stride].
	// If RowMaj is true, then element (i, j) is at Elems[i*Stride + j].
	Elems []float64
}

func NewMat(rows, cols int) *Mat {
	return &Mat{
		Rows:   rows,
		Cols:   cols,
		Stride: rows,
		Elems:  make([]float64, rows*cols),
	}
}

func (a *Mat) Dims() (rows, cols int) {
	return a.Rows, a.Cols
}

func (a *Mat) At(i, j int) float64 {
	if a.RowMaj {
		return a.Elems[i*a.Stride+j]
	}
	return a.Elems[i+j*a.Stride]
}

func (a *Mat) Set(i, j int, x float64) {
	if a.RowMaj {
		a.Elems[i*a.Stride+j] = x
	}
	a.Elems[i+j*a.Stride] = x
}

// T is a non-copying transpose.
func (a *Mat) T() *Mat {
	return &Mat{
		Rows:   a.Cols,
		Cols:   a.Rows,
		Stride: a.Stride,
		RowMaj: !a.RowMaj,
		Elems:  a.Elems,
	}
}

type Transpose byte

const (
	NoTrans = Transpose('N')
	Trans   = Transpose('T')
)

func trans(a *Mat) Transpose {
	if a.RowMaj {
		return Trans
	}
	return NoTrans
}
