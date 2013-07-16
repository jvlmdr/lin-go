package lapack

// ZGELSD functions which use CGo.

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// Called by SolveComplexXxx.
// http://www.netlib.org/lapack/complex16/zgelsd.f
func ZGELSD(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64, work []complex128, lwork int, rwork []float64, iwork []C.integer) (rank int, info int) {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublecomplex
		lda_   = C.integer(lda)
		b_     *C.doublecomplex
		ldb_   = C.integer(ldb)
		s_     *C.doublereal
		rcond_ = C.doublereal(rcond)
		rank_  C.integer
		work_  *C.doublecomplex
		lwork_ = C.integer(lwork)
		rwork_ *C.doublereal
		iwork_ *C.integer
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublecomplex)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublecomplex)(unsafe.Pointer(&b[0]))
	}
	if len(s) > 0 {
		s_ = (*C.doublereal)(unsafe.Pointer(&s[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublecomplex)(unsafe.Pointer(&work[0]))
	}
	if len(rwork) > 0 {
		rwork_ = (*C.doublereal)(unsafe.Pointer(&rwork[0]))
	}
	if len(iwork) > 0 {
		iwork_ = &iwork[0]
	}

	C.zgelsd_(&m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, s_, &rcond_, &rank_,
		work_, &lwork_, rwork_, iwork_, &info_)

	return int(rank_), int(info_)
}
