package mat

// Returns a thin-wrapper which vectorizes a constant matrix.
func VecColMajor(A Const) ConstAsVec {
	return ConstAsVec{A, false}
}

// Returns a thin-wrapper which vectorizes a constant matrix.
func VecRowMajor(A Const) ConstAsVec {
	return ConstAsVec{A, true}
}

// Thin-wrapper which vectorizes a constant matrix.
type ConstAsVec struct {
	Matrix   Const
	RowMajor bool
}

// Returns a thin-wrapper which vectorizes a mutable matrix.
func MutableVecColMajor(A Mutable) MutableAsVec {
	return MutableAsVec{A, false}
}

// Returns a thin-wrapper which vectorizes a mutable matrix.
func MutableVecColMajor(A Mutable) MutableAsVec {
	return MutableAsVec{A, true}
}

// Thin-wrapper which vectorizes a mutable matrix.
type MutableAsVec struct {
	Matrix   Mutable
	RowMajor bool
}
