package lapack

// zgels functions which use CGo and are trivially mapped to zgels functions.

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

func zgels(trans Transpose, m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, work []complex128, lwork int) int {
	var (
		trans_ = C.char(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublecomplex
		lda_   = C.integer(lda)
		b_     *C.doublecomplex
		ldb_   = C.integer(ldb)
		work_  *C.doublecomplex
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublecomplex)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublecomplex)(unsafe.Pointer(&b[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	}

	C.zgels_(&trans_, &m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, work_, &lwork_, &info_)
	return int(info_)
}
