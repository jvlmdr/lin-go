package lapack

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// http://www.netlib.org/lapack/double/zsytrf.f
func zsytrf(uplo UpLo, n int, a []complex128, lda int, ipiv IntList, work []complex128, lwork int) int {
	var (
		uplo_  = C.char(uplo)
		n_     = C.integer(n)
		a_     *C.doublecomplex
		lda_   = C.integer(lda)
		ipiv_  *C.integer
		work_  *C.doublecomplex
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublecomplex)(unsafe.Pointer(&a[0]))
	}
	if len(ipiv) > 0 {
		ipiv_ = &ipiv[0]
	}
	if len(work) > 0 {
		work_ = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	}

	C.zsytrf_(&uplo_, &n_, a_, &lda_, ipiv_, work_, &lwork_, &info_)
	return int(info_)
}

// http://www.netlib.org/lapack/double/zsytrs.f
func zsytrs(uplo UpLo, n, nrhs int, a []complex128, lda int, ipiv IntList, b []complex128, ldb int) int {
	var (
		uplo_ = C.char(uplo)
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

	C.zsytrs_(&uplo_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)
	return int(info_)
}
