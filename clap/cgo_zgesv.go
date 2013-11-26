package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGESV: complex double-precision GEneral SolVe
//
// http://www.netlib.org/lapack/complex16/zgesv.f
func zgesv(n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	ipiv := make([]C.integer, n)
	return zgesvHelper(n, nrhs, a, lda, ipiv, b, ldb)
}

func zgesvHelper(n, nrhs int, a []complex128, lda int, ipiv []C.integer, b []complex128, ldb int) error {
	var (
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
		ipiv_ = ptrInt(ipiv)
		b_    = ptrComplex128(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.zgesv_(&n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)
	return zgetrfError(int(info_))
}
