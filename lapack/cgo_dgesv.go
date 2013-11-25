package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGESV: (Double-precision) GEneral SolVe
//
// http://www.netlib.org/lapack/double/dgesv.f
func dgesv(n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	ipiv := make([]C.integer, n)
	return dgesvHelper(n, nrhs, a, lda, ipiv, b, ldb)
}

func dgesvHelper(n, nrhs int, a []float64, lda int, ipiv []C.integer, b []float64, ldb int) error {
	defer runtime.GC()

	var (
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
		ipiv_ = ptrInt(ipiv)
		b_    = ptrFloat64(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.dgesv_(&n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)
	return dgetrfError(int(info_))
}
