package lapack

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// http://www.netlib.org/lapack/double/dgeqrf.f
func dgeqrf(m, n int, a []float64, lda int, tau, work []float64, lwork int) int {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		tau_   *C.doublereal
		work_  *C.doublereal
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(tau) > 0 {
		tau_ = (*C.doublereal)(unsafe.Pointer(&tau[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}

	C.dgeqrf_(&m_, &n_, a_, &lda_, tau_, work_, &lwork_, &info_)
	return int(info_)
}

// http://www.netlib.org/lapack/double/dormqr.f
func dormqr(side Side, trans Transpose, m, n, k int, a []float64, lda int, tau []float64, c []float64, ldc int, work []float64, lwork int) int {
	var (
		side_  = C.char(side)
		trans_ = C.char(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		k_     = C.integer(k)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		tau_   *C.doublereal
		c_     *C.doublereal
		ldc_   = C.integer(ldc)
		work_  *C.doublereal
		lwork_ = C.integer(lwork)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(tau) > 0 {
		tau_ = (*C.doublereal)(unsafe.Pointer(&tau[0]))
	}
	if len(c) > 0 {
		c_ = (*C.doublereal)(unsafe.Pointer(&c[0]))
	}
	if len(work) > 0 {
		work_ = (*C.doublereal)(unsafe.Pointer(&work[0]))
	}

	C.dormqr_(&side_, &trans_, &m_, &n_, &k_, a_, &lda_, tau_, c_, &ldc_, work_, &lwork_, &info_)
	return int(info_)
}

// http://www.netlib.org/lapack/double/dtrtrs.f
func dtrtrs(uplo UpLo, trans Transpose, diag Diag, n, nrhs int, a []float64, lda int, b []float64, ldb int) int {
	var (
		uplo_  = C.char(uplo)
		trans_ = C.char(trans)
		diag_  = C.char(diag)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		b_     *C.doublereal
		ldb_   = C.integer(ldb)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}

	C.dtrtrs_(&uplo_, &trans_, &diag_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)
	return int(info_)
}
