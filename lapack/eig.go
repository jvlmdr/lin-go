package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"

	"fmt"
)

func EigSymm(A mat.Stride) ([]float64, error) {
	if !A.Size().Square() {
		panic(fmt.Sprint("A is not square: ", A.Size()))
	}

	n := mat.Rows(A)
	v := make([]float64, n)
	info := dsyevAuto(Values, UpperTriangle, n, A.Elems, A.Stride, v)
	if info != 0 {
		return nil, ErrNonZeroInfo(info)
	}
	return v, nil
}

func dsyevAuto(jobz Jobz, uplo UpLo, n int, A []float64, lda int, w []float64) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info := dsyev(jobz, uplo, n, A, lda, w, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dsyev(jobz, uplo, n, A, lda, w, work, lwork)
}
