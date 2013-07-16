package blas

// #include "../f2c.h"
// #include "../clapack.h"
import "C"

// http://www.netlib.org/blas/dgemv.f
func DGEMV(trans Transpose, m, n int, alpha float64, A []float64, lda int, x []float64, incx int, beta float64, y []float64, incy int) {
	var (
		trans_ = C.char(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		alpha_ = C.doublereal(alpha)
		A_     = (*C.doublereal)(&A[0])
		lda_   = C.integer(lda)
		x_     = (*C.doublereal)(&x[0])
		incx_  = C.integer(incx)
		beta_  = C.doublereal(beta)
		y_     = (*C.doublereal)(&y[0])
		incy_  = C.integer(incy)
	)

	C.dgemv_(&trans_, &m_, &n_, &alpha_, A_, &lda_, x_, &incx_, &beta_, y_, &incy_)
}
