package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DSYTRS: (Double-precision) SYmmetric TRiangular factor Solve
//
// http://www.netlib.org/lapack/double/dsytrs.f
func dsytrs(uplo Triangle, n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int) error {
	return dsytrsHelper(uplo, n, nrhs, a, lda, toCInt(ipiv), b, ldb)
}

func dsytrsHelper(uplo Triangle, n, nrhs int, a []float64, lda int, ipiv []C.integer, b []float64, ldb int) error {
	defer runtime.GC()

	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
		ipiv_ = ptrInt(ipiv)
		b_    = ptrFloat64(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.dsytrs_(&uplo_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)

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
