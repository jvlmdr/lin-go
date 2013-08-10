package lapack

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// Called by SolveComplexSquareXxx.
// http://www.netlib.org/lapack/double/zgesv.f
func zgesv(n, nrhs int, a []complex128, lda int, ipiv IntList, b []complex128, ldb int) int {
	var (
		n_    = C.integer(n)
		nrhs_ = C.integer(nrhs)
		a_    *C.doublecomplex
		lda_  = C.integer(lda)
		ipiv_ *C.integer
		b_    *C.doublecomplex
		ldb_  = C.integer(ldb)
		info_ C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublecomplex)(unsafe.Pointer(&a[0]))
	}
	if len(ipiv) > 0 {
		ipiv_ = &ipiv[0]
	}
	if len(b) > 0 {
		b_ = (*C.doublecomplex)(unsafe.Pointer(&b[0]))
	}

	C.zgesv_(&n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)
	return int(info_)
}
