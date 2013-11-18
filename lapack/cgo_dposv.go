package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DPOSV: (Double-precision) POsitive-definite SolVe
//
// http://www.netlib.org/lapack/double/dposv.f
func dposv(n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	var (
		uplo_ = C.char(DefaultTri)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
		b_    = ptrFloat64(b)
		ldb_  = C.integer(ldb)
	)
	var info_ C.integer

	C.dposv_(&uplo_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)
	return dpotrfError(int(info_))
}
