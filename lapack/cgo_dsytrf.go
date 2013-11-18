package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DSYTRF: (Double-precision) SYmmetric TRiangular Factor
//
// http://www.netlib.org/lapack/double/dsytrf.f
func dsytrf(uplo Triangle, n int, a []float64, lda int) (ipiv []int, err error) {
	ipiv_ := make([]C.integer, n)

	// Query workspace size.
	work := make([]float64, 1)
	err = dsytrfHelper(uplo, n, a, lda, ipiv_, work, -1)
	if err != nil {
		return nil, err
	}

	lwork := int(work[0])
	work = make([]float64, lwork)
	err = dsytrfHelper(uplo, n, a, lda, ipiv_, work, lwork)
	if err != nil {
		return nil, err
	}
	return fromCInt(ipiv_), nil
}

func dsytrfHelper(uplo Triangle, n int, a []float64, lda int, ipiv []C.integer, work []float64, lwork int) error {
	var (
		uplo_  = C.char(uplo)
		n_     = C.integer(n)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dsytrf_(&uplo_, &n_, a_, &lda_, ipiv_, work_, &lwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errSingular(info)
	default:
		return nil
	}
}
