package cmat

// Describes a matrix with a size.
type dimser interface {
	Dims() (int, int)
}

// Describes a read-only matrix.
type Const interface {
	Dims() (int, int)
	At(i, j int) complex128
}

// Describes a fixed-size matrix whose elements can be modified.
type Mutable interface {
	Const
	Set(i, j int, x complex128)
}
