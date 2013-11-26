package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGETRS: complex double-precision GEneral TRiangular factor Solve
//
// http://www.netlib.org/lapack/complex16/zgetrs.f
func zgetrs(h bool, n, nrhs int, a []complex128, lda int, ipiv []int, b []complex128, ldb int) error {
	return zgetrsHelper(h, n, nrhs, a, lda, toCInt(ipiv), b, ldb)
}

func zgetrsHelper(h bool, n, nrhs int, a []complex128, lda int, ipiv []C.integer, b []complex128, ldb int) error {
	var (
		trans_ = conjTransChar(h)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		b_     = ptrComplex128(b)
		ldb_   = C.integer(ldb)
	)
	var info_ C.integer

	C.zgetrs_(&trans_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)

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
