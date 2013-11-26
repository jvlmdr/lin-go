package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGEQRF: complex double-precision GEneral QR Factor
//
// http://www.netlib.org/lapack/complex16/zgeqrf.f
func zgeqrf(m, n int, a []complex128, lda int, tau []complex128) error {
	// Query workspace size.
	work := make([]complex128, 1)
	err := zgeqrfHelper(m, n, a, lda, tau, work, -1)
	if err != nil {
		return err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	return zgeqrfHelper(m, n, a, lda, tau, work, lwork)
}

func zgeqrfHelper(m, n int, a []complex128, lda int, tau, work []complex128, lwork int) error {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		tau_   = ptrComplex128(tau)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.zgeqrf_(&m_, &n_, a_, &lda_, tau_, work_, &lwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info == 0:
		return nil
	default:
		panic(errUnknown(info))
	}
}
