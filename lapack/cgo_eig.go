package lapack

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

func dsyev(jobz Jobz, uplo UpLo, n int, a []float64, lda int, w, work []float64, lwork int) int {
	var (
		jobz_  = C.char(jobz)
		uplo_  = C.char(uplo)
		n_     = C.integer(n)
		a_     = nonEmptyPtrFloat64(a)
		lda_   = C.integer(lda)
		w_     = nonEmptyPtrFloat64(w)
		work_  = nonEmptyPtrFloat64(work)
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	C.dsyev_(&jobz_, &uplo_, &n_, a_, &lda_, w_, work_, &lwork_, &info_)
	return int(info_)
}
