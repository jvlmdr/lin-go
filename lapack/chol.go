package lapack

import "github.com/jackvalmadre/lin-go/mat"

type CholFact struct {
	UpLo UpLo
	A    mat.ColMajor
}

func Chol(A mat.ColMajor) (CholFact, error) {
	// Check that A is square.
	if !A.Size().Square() {
		panic("matrix is not square")
	}

	const uplo = LowerTriangle

	a := A.ColMajorArray()
	info := dpotrf(uplo, mat.Rows(A), a, A.ColStride())
	if info != 0 {
		return CholFact{}, ErrNonZeroInfo
	}
	ldl := CholFact{uplo, A}
	return ldl, nil
}

func (f CholFact) Solve(B mat.ColMajor) error {
	// Check that B has the same number of rows as A.
	if mat.Rows(f.A) != mat.Rows(B) {
		panic("numbers of rows do not match")
	}

	a := f.A.ColMajorArray()
	b := B.ColMajorArray()
	info := dpotrs(f.UpLo, mat.Rows(f.A), mat.Cols(B), a, f.A.ColStride(), b, B.ColStride())
	if info != 0 {
		return ErrNonZeroInfo
	}
	return nil
}
