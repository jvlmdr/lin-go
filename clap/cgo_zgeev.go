package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGEEV: complex double-precision GEneral EigenValues
//
// http://www.netlib.org/lapack/complex16/zgeev.f
func zgeev(jobvl, jobvr jobzMode, n int, a []complex128, lda int, vl []complex128, ldvl int, vr []complex128, ldvr int) (w []complex128, err error) {
	w = make([]complex128, n)
	rwork := make([]float64, 2*n)

	// Query workspace size.
	work := make([]complex128, 1)
	err = zgeevHelper(jobvl, jobvr, n, a, lda, w, vl, ldvl, vr, ldvr, work, -1, rwork)
	if err != nil {
		return nil, err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	err = zgeevHelper(jobvl, jobvr, n, a, lda, w, vl, ldvl, vr, ldvr, work, lwork, rwork)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func zgeevHelper(jobvl, jobvr jobzMode, n int, a []complex128, lda int, w, vl []complex128, ldvl int, vr []complex128, ldvr int, work []complex128, lwork int, rwork []float64) error {
	var (
		jobvl_ = jobzChar(jobvl)
		jobvr_ = jobzChar(jobvr)
		n_     = C.integer(n)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		w_     = ptrComplex128(w)
		vl_    = ptrComplex128(vl)
		ldvl_  = C.integer(ldvl)
		vr_    = ptrComplex128(vr)
		ldvr_  = C.integer(ldvr)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
		rwork_ = ptrFloat64(rwork)
	)
	var info_ C.integer

	C.zgeev_(&jobvl_, &jobvr_, &n_, a_, &lda_, w_, vl_, &ldvl_, vr_, &ldvr_, work_, &lwork_, rwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errOffDiagFailConverge(info)
	default:
		return nil
	}
}
