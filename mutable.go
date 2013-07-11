package mat

import "github.com/jackvalmadre/go-vec"

// Copies from a Const matrix to a Mutable matrix.
// The size of A must match that of B.
func Copy(A Mutable, B Const) {
	if !A.Size().Equals(B.Size()) {
		panic(ErrNotSameSize)
	}
	vec.Copy(MutableVec(A), Vec(B))
}
