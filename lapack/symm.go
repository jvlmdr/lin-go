package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
)

// Vector version of SolveSymmMat().
// Like SolveSymmNoCopy(), except A and b are left intact.
func SolveSymm(A mat.Const, b vec.Const) (vec.Slice, error) {
	return SolveSymmNoCopy(mat.MakeStrideCopy(A), vec.MakeSliceCopy(b))
}

// Like SolveSymmMatNoCopy(), except A and B are left intact.
func SolveSymmMat(A mat.Const, B mat.Const) (mat.Stride, error) {
	return SolveSymmMatNoCopy(mat.MakeStrideCopy(A), mat.MakeStrideCopy(B))
}

// Vector version of SolveSymmMat.
func SolveSymmNoCopy(A mat.Stride, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	X, err := SolveSymmMatNoCopy(A, B)
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is symmetric.
//
// Calls dsysv.
//
// Overwrites A and B.
// Returns X which references the elements of B.
func SolveSymmMatNoCopy(A mat.Stride, B mat.Stride) (mat.Stride, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic(fmt.Sprintf("matrix is not square: %v", A.Size()))
	}
	// Check that B has the same number of rows as A.
	if A.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", A.Size(), B.Size()))
	}

	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, A.Rows)

	info := dsysvAuto(uplo, A.Rows, B.Cols, A.Elems, A.Stride, ipiv, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo
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
