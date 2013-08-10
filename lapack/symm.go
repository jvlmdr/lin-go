package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Vector version of SolveNSymm.
func SolveSymm(A mat.ColMajor, b vec.Slice) (vec.Slice, error) {
	B := mat.FromSlice(b)
	if _, err := SolveNSymm(A, B); err != nil {
		return vec.Slice{}, err
	}
	return b, nil
}

// Solves A X = B where A is symmetric.
//
// Calls dsysv.
//
// Overwrites A and B.
// Returns X which references the elements of B.
func SolveNSymm(A mat.ColMajor, B mat.ColMajor) (mat.ColMajor, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic("matrix is not square")
	}
	// Check that B has the same number of rows as A.
	if mat.Rows(A) != mat.Rows(B) {
		panic("numbers of rows do not match")
	}

	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, mat.Rows(A))

	a := A.ColMajorArray()
	b := B.ColMajorArray()
	info := dsysvAuto(uplo, mat.Rows(A), mat.Cols(B), a, A.Stride(), ipiv, b, B.Stride())
	if info != 0 {
		return nil, ErrNonZeroInfo
	}
	return B, nil
}

// Automatically allocates workspace.
func dsysvAuto(uplo UpLo, n, nrhs int, a []float64, lda int, ipiv IntList, b []float64, ldb int) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	if info := dsysv(uplo, n, nrhs, a, lda, ipiv, b, ldb, work, lwork); info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dsysv(uplo, n, nrhs, a, lda, ipiv, b, ldb, work, lwork)
}
