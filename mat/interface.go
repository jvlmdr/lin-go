package mat

// Describes a read-only matrix.
type Const interface {
	Dims() (int, int)
	At(i, j int) float64
}

// Describes a fixed-size matrix whose elements can be modified.
type Mutable interface {
	Const
	Set(i, j int, x float64)
}
