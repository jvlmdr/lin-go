package zmat

// Returns MutableH(A).
func (A SubContiguous) H() Mutable { return MutableH(A) }
