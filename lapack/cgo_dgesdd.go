package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGESDD: (Double-precision) GEneral SvD by Divide-and-conquer
//
// http://www.netlib.org/lapack/double/dgesdd.f
func dgesdd(m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int) error {
	var err error
	// Query workspace size.
	work := make([]float64, 1)
	iwork := make([]C.integer, 1)
	err = dgesddHelper(m, n, a, lda, s, u, ldu, vt, ldvt, work, -1, iwork)
	if err != nil {
		return err
	}

	lwork := int(work[0])
	liwork := int(iwork[0])
	work = make([]float64, lwork)
	iwork = make([]C.integer, liwork)

	return dgesddHelper(m, n, a, lda, s, u, ldu, vt, ldvt, work, lwork, iwork)
}

func dgesddHelper(m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int, iwork []C.integer) error {
	defer runtime.GC()

	var (
		jobz_  = C.char('S')
		m_     = C.integer(m)
		n_     = C.integer(n)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		s_     = ptrFloat64(s)
		u_     = ptrFloat64(u)
		ldu_   = C.integer(ldu)
		vt_    = ptrFloat64(vt)
		ldvt_  = C.integer(ldvt)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
		iwork_ = ptrInt(iwork)
	)
	var info_ C.integer

	C.dgesdd_(&jobz_, &m_, &n_, a_, &lda_, s_, u_, &ldu_, vt_, &ldvt_, work_, &lwork_, iwork_, &info_)

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
