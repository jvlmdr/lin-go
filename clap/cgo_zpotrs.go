package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZPOTRS: complex double-precision POsitive-definite TRiangular factor Solve
//
// http://www.netlib.org/lapack/complex16/zpotrs.f
func zpotrs(uplo Triangle, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
		b_    = ptrComplex128(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.zpotrs_(&uplo_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)

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
