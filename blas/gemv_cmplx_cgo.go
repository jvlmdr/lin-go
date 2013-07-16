package blas

import "unsafe"

// #include "../f2c.h"
// #include "../blas.h"
import "C"

// http://www.netlib.org/blas/zgemv.f
func ZGEMV(trans Transpose, m, n int, alpha complex128, A []complex128, lda int, x []complex128, incx int, beta complex128, y []complex128, incy int) {
	var (
		trans_ = C.char(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		alpha_ = (*C.doublecomplex)(unsafe.Pointer(&alpha))
		A_     = (*C.doublecomplex)(unsafe.Pointer(&A[0]))
		lda_   = C.integer(lda)
		x_     = (*C.doublecomplex)(unsafe.Pointer(&x[0]))
		incx_  = C.integer(incx)
		beta_  = (*C.doublecomplex)(unsafe.Pointer(&beta))
		y_     = (*C.doublecomplex)(unsafe.Pointer(&y[0]))
		incy_  = C.integer(incy)
	)

	C.zgemv_(&trans_, &m_, &n_, alpha_, A_, &lda_, x_, &incx_, beta_, y_, &incy_)
}
