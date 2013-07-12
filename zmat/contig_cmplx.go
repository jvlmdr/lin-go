package zmat

// Returns MutableH(A).
func (A Contiguous) H() Mutable { return MutableH(A) }
