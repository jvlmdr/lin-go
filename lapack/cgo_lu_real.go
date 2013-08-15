package lapack

import "unsafe"

// #include "../f2c.h"
// #include "../lapack.h"
import "C"

// http://www.netlib.org/lapack/double/dgetrf.f
func dgetrf(m, n int, a []float64, lda int, ipiv IntList) int {
	var (
		m_    = C.integer(m)
		n_    = C.integer(n)
		a_    *C.doublereal
		lda_  = C.integer(lda)
		ipiv_ *C.integer
		info_ C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(ipiv) > 0 {
		ipiv_ = &ipiv[0]
	}

	C.dgetrf_(&m_, &n_, a_, &lda_, ipiv_, &info_)
	return int(info_)
}

// http://www.netlib.org/lapack/double/dgetrs.f
func dgetrs(trans Transpose, n, nrhs int, a []float64, lda int, ipiv IntList, b []float64, ldb int) int {
	var (
		trans_ = C.char(trans)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     *C.doublereal
		lda_   = C.integer(lda)
		ipiv_  *C.integer
		b_     *C.doublereal
		ldb_   = C.integer(ldb)
		info_  C.integer
	)

	if len(a) > 0 {
		a_ = (*C.doublereal)(unsafe.Pointer(&a[0]))
	}
	if len(ipiv) > 0 {
		ipiv_ = &ipiv[0]
	}
	if len(b) > 0 {
		b_ = (*C.doublereal)(unsafe.Pointer(&b[0]))
	}

	C.dgetrs_(&trans_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, &info_)
	return int(info_)
}
