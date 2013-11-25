package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DPOTRS: (Double-precision) POsitive-definite TRiangular factor Solve
//
// http://www.netlib.org/lapack/double/dpotrs.f
func dpotrs(uplo Triangle, n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	defer runtime.GC()

	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
		b_    = ptrFloat64(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.dpotrs_(&uplo_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)

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
