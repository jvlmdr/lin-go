package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGETRS: (Double-precision) GEneral TRiangular factor Solve
//
// http://www.netlib.org/lapack/double/dgetrs.f
func dgetrs(trans bool, n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int) error {
	return dgetrsHelper(trans, n, nrhs, a, lda, toCInt(ipiv), b, ldb)
}

func dgetrsHelper(trans bool, n, nrhs int, a []float64, lda int, ipiv []C.integer, b []float64, ldb int) error {
	var (
		trans_ = transChar(trans)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		b_     = ptrFloat64(b)
		ldb_   = C.integer(ldb)
	)
	var info_ C.integer

	C.dgetrs_(&trans_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info == 0:
		return nil
	default:
		panic(errUnknown(info))
	}
}
