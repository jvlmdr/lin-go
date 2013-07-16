package lapack

// DGELS functions which use CGo and are trivially mapped to ZGELS functions.

import "unsafe"

// #include "../f2c.h"
// #include "../clapack.h"
import "C"

// Called by SolveFullRankXxx.
func DGELS(trans Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) int {
	var (
		trans_ = C.char(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		b_     *C.doublereal
		ldb_   = C.integer(ldb)
		work_  *C.doublereal
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}

	C.dgels_(&trans_, &m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, work_, &lwork_, &info_)
	return int(info_)
}
