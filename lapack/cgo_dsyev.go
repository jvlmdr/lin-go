package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DSYEV: (Double-precision) SYmmetric EigenValues
//
// http://www.netlib.org/lapack/double/dsyev.f
func dsyev(jobz jobzMode, uplo Triangle, n int, a []float64, lda int) ([]float64, error) {
	var err error
	w := make([]float64, n)

	// Query workspace size.
	work := make([]float64, 1)
	err = dsyevHelper(jobz, uplo, n, a, lda, w, work, -1)
	if err != nil {
		return nil, err
	}

	lwork := int(work[0])
	work = make([]float64, max(1, lwork))
	err = dsyevHelper(jobz, uplo, n, a, lda, w, work, lwork)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func dsyevHelper(jobz jobzMode, uplo Triangle, n int, a []float64, lda int, w, work []float64, lwork int) error {
	var (
		jobz_  = jobzChar(jobz)
		uplo_  = uploChar(uplo)
		n_     = C.integer(n)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		w_     = ptrFloat64(w)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dsyev_(&jobz_, &uplo_, &n_, a_, &lda_, w_, work_, &lwork_, &info_)

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
