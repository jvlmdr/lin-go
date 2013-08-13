package lapack

// dgels functions which are trivially mapped to ZGELS functions.

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
)

// Solves A x = b where A is full rank.
//
// Calls dgels.
func SolveFullRank(A mat.Const, b vec.Const) vec.Slice {
	if mat.Rows(A) != b.Len() {
		panic("Number of equations does not match dimension of vector")
	}

	// Translate A x = b into Q x = u.
	m, n := mat.RowsCols(A)
	Q := mat.MakeContigCopy(A)
	// Allocate enough space for input and solution.
	ux := vec.MakeSlice(max(m, n))
	u := ux.Subvec(0, m)
	vec.Copy(u, b)

	return SolveFullRankInPlace(Q, NoTrans, ux)
}

// Solves A x = b where A is full rank.
//
// Calls dgels.
func SolveFullRankInPlace(A mat.ColMajor, trans Transpose, b vec.Slice) vec.Slice {
	B := mat.Contig{b.Len(), 1, []float64(b)}
	SolveNFullRankInPlace(A, trans, B)
	//X := SolveNFullRankInPlace(A, trans, B)
	return b[0:mat.Cols(A)]
}

// Solves A X = B where A is full rank.
//
// Calls dgels.
func SolveNFullRank(A mat.Const, B mat.Const) mat.ColMajor {
	if mat.Rows(A) != mat.Rows(B) {
		panic("Matrices have different number of rows")
	}

	// Translate into Q X = U.
	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)
	Q := mat.MakeContigCopy(A)
	// Allocate enough space for constraints and solution.
	UX := mat.MakeContig(max(m, n), nrhs)
	U := UX.Submat(mat.MakeRect(0, 0, m, nrhs))
	mat.Copy(U, B)
	return SolveNFullRankInPlace(Q, NoTrans, UX)
}

// Solves A X = B where A is full rank.
//
// Calls dgels.
//
// B must be large enough to hold both the constraints and the solution (not simultaneously).
// Returns a matrix which references the elements of B.
// A will be over-written with either the LQ or QR factorization.
func SolveNFullRankInPlace(A mat.ColMajor, trans Transpose, B mat.ColMajor) mat.ColMajor {
	size := A.Size()
	// Transpose dimensions if necessary.
	if trans != NoTrans {
		size = size.T()
	}
	// Check that B has enough space to contain input and solution.
	if mat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if mat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	dgelsAuto(trans, mat.Rows(A), mat.Cols(A), mat.Cols(B),
		A.ColMajorArray(), A.ColStride(), B.ColMajorArray(), B.ColStride())

	return mat.Stride{mat.Cols(A), mat.Cols(B), B.ColStride(), B.ColMajorArray()}
}

// Automatically allocates workspace.
func dgelsAuto(trans Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int) (info int) {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info = dgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
	if info != 0 {
		return
	}

	lwork = int(forceToReal(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
}
