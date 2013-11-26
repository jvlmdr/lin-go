package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGESDD: complex double-precision GEneral SvD by Divide-and-conquer
//
// http://www.netlib.org/lapack/complex16/zgesdd.f
func zgesdd(m, n int, a []complex128, lda int, s []float64, u []complex128, ldu int, vt []complex128, ldvt int) error {
	var err error

	lrwork := min(m, n) * max(5*min(m, n)+7, 2*max(m, n)+2*min(m, n)+1)
	rwork := make([]float64, lrwork)
	liwork := 8 * min(m, n)
	iwork := make([]C.integer, liwork)

	// Query workspace size.
	work := make([]complex128, 1)
	err = zgesddHelper(m, n, a, lda, s, u, ldu, vt, ldvt, work, -1, rwork, iwork)
	if err != nil {
		return err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))

	return zgesddHelper(m, n, a, lda, s, u, ldu, vt, ldvt, work, lwork, rwork, iwork)
}

func zgesddHelper(m, n int, a []complex128, lda int, s []float64, u []complex128, ldu int, vt []complex128, ldvt int, work []complex128, lwork int, rwork []float64, iwork []C.integer) error {
	var (
		jobz_  = C.char('S')
		m_     = C.integer(m)
		n_     = C.integer(n)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		s_     = ptrFloat64(s)
		u_     = ptrComplex128(u)
		ldu_   = C.integer(ldu)
		vt_    = ptrComplex128(vt)
		ldvt_  = C.integer(ldvt)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
		rwork_ = ptrFloat64(rwork)
		iwork_ = ptrInt(iwork)
	)
	var info_ C.integer

	C.zgesdd_(&jobz_, &m_, &n_, a_, &lda_, s_, u_, &ldu_, vt_, &ldvt_, work_, &lwork_, rwork_, iwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errFailConverge(info)
	default:
		return nil
	}
}
