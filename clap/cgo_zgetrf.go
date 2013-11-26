package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGETRF: complex double-precision GEneral TRiangular Factor
//
// http://www.netlib.org/lapack/complex16/zgetrf.f
func zgetrf(m, n int, a []complex128, lda int) (ipiv []int, err error) {
	ipiv_ := make([]C.integer, min(m, n))
	err = zgetrfHelper(m, n, a, lda, ipiv_)
	if err != nil {
		return nil, err
	}
	return fromCInt(ipiv_), nil
}

func zgetrfHelper(m, n int, a []complex128, lda int, ipiv []C.integer) error {
	var (
		m_    = C.integer(m)
		n_    = C.integer(n)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
		ipiv_ = ptrInt(ipiv)
	)
	var info_ C.integer

	C.zgetrf_(&m_, &n_, a_, &lda_, ipiv_, &info_)
	return zgetrfError(int(info_))
}

func zgetrfError(info int) error {
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errSingular(info)
	default:
		return nil
	}
}
