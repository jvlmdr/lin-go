package lapack

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// http://www.netlib.org/lapack/double/dposv.f
func dposv(uplo UpLo, n, nrhs int, a []float64, lda int, b []float64, ldb int) int {
	var (
		uplo_ = C.char(uplo)
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    *C.doublereal
		lda_  = C.integer(lda)
		b_    *C.doublereal
		ldb_  = C.integer(ldb)
		info_ C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}

	C.dposv_(&uplo_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)
	return int(info_)
}
