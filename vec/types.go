// A vector package without matrices.

package vec

// Describes a list of real numbers.
type Const interface {
	// Returns the dimension of the vector space.
	Len() int
	// Accesses the i-th element.
	At(i int) float64
}

// Describes a list of real numbers which can be modified.
type Mutable interface {
	Const
	// Modifies the i-th element.
	Set(i int, x float64)
}

// Knows the dimension of the vector space.
type Space interface {
	// Returns the dimension of the vector space.
	Len() int
}

// Knows the dimension of the vector space and how to create a new vector.
type Type interface {
	Space
	// Creates a zero vector with the same type.
	New() MutableTyped
}

type Typed interface {
	Type() Type
}

// Describes a list of real numbers.
type ConstTyped interface {
	Const
	Typed
}

// Describes a list of real numbers which can be modified.
type MutableTyped interface {
	Mutable
	Typed
}
