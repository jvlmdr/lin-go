package mat

// Default (column-major contiguous) matrix.
func Make(rows, cols int) Mutable {
	return MakeColMajorContiguous(rows, cols)
}

// Default (column-major) contiguous matrix.
func MakeContiguous(rows, cols int) Contiguous {
	return MakeColMajorContiguous(rows, cols)
}

// Default (column-major) vectorization.
func Vec(A Const) {
	return VecColMajor(A)
}

// Default (column-major) vectorization.
func MutableVec(A Mutable) MutableAsVec {
	return MutableVecColMajor(A)
}

// Default ordering for matricization is column-major.
func Mat(x vec.Const, rows, cols int) ConstMatExpr {
	return
}

// Default ordering.
func MutableMat(x vec.Mutable, rows, cols int) MutableMatExpr {
}
