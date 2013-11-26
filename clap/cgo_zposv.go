package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZPOSV: complex double-precision POsitive-definite SolVe
//
// http://www.netlib.org/lapack/complex16/zposv.f
func zposv(n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	var (
		uplo_ = C.char(DefaultTri)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
		b_    = ptrComplex128(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.zposv_(&uplo_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)
	return zpotrfError(int(info_))
}
