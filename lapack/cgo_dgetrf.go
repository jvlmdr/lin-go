package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGETRF: (Double-precision) GEneral TRiangular Factor
//
// http://www.netlib.org/lapack/double/dgetrf.f
func dgetrf(m, n int, a []float64, lda int) (ipiv []int, err error) {
	ipiv_ := make([]C.integer, min(m, n))
	err = dgetrfHelper(m, n, a, lda, ipiv_)
	if err != nil {
		return nil, err
	}
	return fromCInt(ipiv_), nil
}

func dgetrfHelper(m, n int, a []float64, lda int, ipiv []C.integer) error {
	var (
		m_    = C.integer(m)
		n_    = C.integer(n)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
		ipiv_ = ptrInt(ipiv)
	)
	var info_ C.integer

	C.dgetrf_(&m_, &n_, a_, &lda_, ipiv_, &info_)
	return dgetrfError(int(info_))
}

func dgetrfError(info int) error {
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errSingular(info)
	default:
		return nil
	}
}
