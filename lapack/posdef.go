package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Vector version of SolveNPosDef.
func SolvePosDef(A mat.ColMajor, b vec.Slice) (vec.Slice, error) {
	B := mat.FromSlice(b)
	if _, err := SolveNPosDef(A, B); err != nil {
		return vec.Slice{}, err
	}
	return b, nil
}

// Solves A X = B where A is symmetric.
//
// Calls dposv.
//
// Overwrites A and B.
// Returns X which references the elements of B.
func SolveNPosDef(A mat.ColMajor, B mat.ColMajor) (mat.ColMajor, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic("matrix is not square")
	}
	// Check that B has the same number of rows as A.
	if mat.Rows(A) != mat.Rows(B) {
		panic("numbers of rows do not match")
	}

	const uplo = LowerTriangle

	a := A.ColMajorArray()
	b := B.ColMajorArray()
	info := dposv(uplo, mat.Rows(A), mat.Cols(B), a, A.Stride(), b, B.Stride())
	if info != 0 {
		return nil, ErrNonZeroInfo
	}
	return B, nil
}
