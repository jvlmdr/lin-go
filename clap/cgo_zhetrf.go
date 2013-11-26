package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZHETRF: complex double-precision HErmitian TRiangular Factor
//
// http://www.netlib.org/lapack/complex16/zhetrf.f
func zhetrf(uplo Triangle, n int, a []complex128, lda int) (ipiv []int, err error) {
	ipiv_ := make([]C.integer, n)

	// Query workspace size.
	work := make([]complex128, 1)
	err = zhetrfHelper(uplo, n, a, lda, ipiv_, work, -1)
	if err != nil {
		return nil, err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	err = zhetrfHelper(uplo, n, a, lda, ipiv_, work, lwork)
	if err != nil {
		return nil, err
	}
	return fromCInt(ipiv_), nil
}

func zhetrfHelper(uplo Triangle, n int, a []complex128, lda int, ipiv []C.integer, work []complex128, lwork int) error {
	var (
		uplo_  = C.char(uplo)
		n_     = C.integer(n)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.zhetrf_(&uplo_, &n_, a_, &lda_, ipiv_, work_, &lwork_, &info_)

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
