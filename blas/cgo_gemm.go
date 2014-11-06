package blas

import "unsafe"

// #include "../f2c.h"
// #include "../blas.h"
import "C"

// http://www.netlib.org/blas/dgemm.f
func dgemm(transa, transb Transpose, m, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	var (
		transa_ = C.char(transa)
		transb_ = C.char(transb)
		m_      = C.integer(m)
		n_      = C.integer(n)
		k_      = C.integer(k)
		alpha_  = (*C.doublereal)(unsafe.Pointer(&alpha))
		a_      *C.doublereal
		lda_    = C.integer(lda)
		b_      *C.doublereal
		ldb_    = C.integer(ldb)
		beta_   = (*C.doublereal)(unsafe.Pointer(&beta))
		c_      *C.doublereal
		ldc_    = C.integer(ldc)
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}
	if len(c) > 0 {
		c_ = (*C.doublereal)(unsafe.Pointer(&c[0]))
	}

	C.dgemm_(&transa_, &transb_, &m_, &n_, &k_, alpha_, a_, &lda_, b_, &ldb_, beta_, c_, &ldc_)
}
