package lapack

import "github.com/jackvalmadre/lin-go/mat"

type LDLFact struct {
	UpLo UpLo
	A    mat.ColMajor
	Ipiv IntList
}

func LDL(A mat.ColMajor) (LDLFact, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic("matrix is not square")
	}

	const uplo = LowerTriangle
	// Permutation indices.
	ipiv := make(IntList, mat.Rows(A))

	a := A.ColMajorArray()
	info := dsytrfAuto(uplo, mat.Rows(A), a, A.ColStride(), ipiv)
	if info != 0 {
		return LDLFact{}, ErrNonZeroInfo
	}
	ldl := LDLFact{uplo, A, ipiv}
	return ldl, nil
}

// Automatically allocates workspace.
func dsytrfAuto(uplo UpLo, n int, a []float64, lda int, ipiv IntList) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	if info := dsytrf(uplo, n, a, lda, ipiv, work, lwork); info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dsytrf(uplo, n, a, lda, ipiv, work, lwork)
}

func (f LDLFact) Solve(B mat.ColMajor) error {
	// Check that B has the same number of rows as A.
	if mat.Rows(f.A) != mat.Rows(B) {
		panic("numbers of rows do not match")
	}

	a := f.A.ColMajorArray()
	b := B.ColMajorArray()
	info := dsytrs(f.UpLo, mat.Rows(f.A), mat.Cols(B), a, f.A.ColStride(), f.Ipiv, b, B.ColStride())
	if info != 0 {
		return ErrNonZeroInfo
	}
	return nil
}
