package lapack

// DGELSD functions which use CGo.

import "unsafe"

// #include "../f2c.h"
// #include "../clapack.h"
import "C"

// Called by SolveXxx.
// http://www.netlib.org/lapack/double/dgelsd.f
func DGELSD(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64, work []float64, lwork int, iwork []C.integer) (rank int, info int) {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		b_     *C.doublereal
		ldb_   = C.integer(ldb)
		s_     *C.doublereal
		rcond_ = C.doublereal(rcond)
		rank_  C.integer
		work_  *C.doublereal
		lwork_ = C.integer(lwork)
		iwork_ *C.integer
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}
	if len(s) > 0 {
		s_ = (*C.doublereal)(unsafe.Pointer(&s[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}
	if len(iwork) > 0 {
		iwork_ = &iwork[0]
	}

	C.dgelsd_(&m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, s_, &rcond_, &rank_,
		work_, &lwork_, iwork_, &info_)

	return int(rank_), int(info_)
}
