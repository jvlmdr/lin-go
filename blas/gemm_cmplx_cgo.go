package blas

import "unsafe"

// #include "../f2c.h"
// #include "../blas.h"
import "C"

// http://www.netlib.org/blas/zgemm.f
func ZGEMM(transa, transb Transpose, m, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	var (
		transa_ = C.char(transa)
		transb_ = C.char(transb)
		m_      = C.integer(m)
		n_      = C.integer(n)
		k_      = C.integer(k)
		alpha_  = (*C.doublecomplex)(unsafe.Pointer(&alpha))
		a_      *C.doublecomplex
		lda_    = C.integer(lda)
		b_      *C.doublecomplex
		ldb_    = C.integer(ldb)
		beta_   = (*C.doublecomplex)(unsafe.Pointer(&beta))
		c_      *C.doublecomplex
		ldc_    = C.integer(ldc)
	)

	if len(a) > 0 {
		a_ = (*C.doublecomplex)(unsafe.Pointer(&a[0]))
	}
	if len(b) > 0 {
		b_ = (*C.doublecomplex)(unsafe.Pointer(&b[0]))
	}
	if len(c) > 0 {
		c_ = (*C.doublecomplex)(unsafe.Pointer(&c[0]))
	}

	C.zgemm_(&transa_, &transb_, &m_, &n_, &k_, alpha_, a_, &lda_, b_, &ldb_, beta_, c_, &ldc_)
}
