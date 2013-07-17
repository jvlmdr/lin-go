// A vector package without matrices.

package zvec

// Describes a list of real numbers.
type Const interface {
	// Returns the vector space which the vector belongs to.
	Len() int
	// Accesses the i-th element.
	At(i int) complex128
}

// Describes a list of real numbers which can be modified.
type Mutable interface {
	Const
	// Modifies the i-th element.
	Set(i int, x complex128)
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
