package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZHEEV: complex double-precision HErmitian EigenValues
//
// http://www.netlib.org/lapack/complex16/zheev.f
func zheev(jobz jobzMode, uplo Triangle, n int, a []complex128, lda int) ([]float64, error) {
	var err error
	w := make([]float64, n)

	lrwork := max(1, 3*n-2)
	rwork := make([]float64, lrwork)

	// Query workspace size.
	work := make([]complex128, 1)
	err = zheevHelper(jobz, uplo, n, a, lda, w, work, -1, rwork)
	if err != nil {
		return nil, err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	err = zheevHelper(jobz, uplo, n, a, lda, w, work, lwork, rwork)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func zheevHelper(jobz jobzMode, uplo Triangle, n int, a []complex128, lda int, w []float64, work []complex128, lwork int, rwork []float64) error {
	var (
		jobz_  = jobzChar(jobz)
		uplo_  = uploChar(uplo)
		n_     = C.integer(n)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		w_     = ptrFloat64(w)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
		rwork_ = ptrFloat64(rwork)
	)
	var info_ C.integer

	C.zheev_(&jobz_, &uplo_, &n_, a_, &lda_, w_, work_, &lwork_, rwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errOffDiagFailConverge(info)
	default:
		return nil
	}
}
