package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZUNMQR: complex double-precision UNitary Multiply by QR
//
// http://www.netlib.org/lapack/complex16/zunmqr.f
func zunmqr(side matSide, h bool, m, n, k int, a []complex128, lda int, tau []complex128, c []complex128, ldc int) error {
	// Query for workspace size.
	work := make([]complex128, 1)
	err := zunmqrHelper(side, h, m, n, k, a, lda, tau, c, ldc, work, -1)
	if err != nil {
		return err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	return zunmqrHelper(side, h, m, n, k, a, lda, tau, c, ldc, work, lwork)
}

func zunmqrHelper(side matSide, h bool, m, n, k int, a []complex128, lda int, tau []complex128, c []complex128, ldc int, work []complex128, lwork int) error {
	var (
		side_  = sideChar(side)
		trans_ = conjTransChar(h)
		m_     = C.integer(m)
		n_     = C.integer(n)
		k_     = C.integer(k)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		tau_   = ptrComplex128(tau)
		c_     = ptrComplex128(c)
		ldc_   = C.integer(ldc)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.zunmqr_(&side_, &trans_, &m_, &n_, &k_, a_, &lda_, tau_, c_, &ldc_, work_, &lwork_, &info_)

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
