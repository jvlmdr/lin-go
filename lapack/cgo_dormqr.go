package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DORMQR: (Double-precision) ORthogonal Multiply by QR
//
// http://www.netlib.org/lapack/double/dormqr.f
func dormqr(side matSide, trans bool, m, n, k int, a []float64, lda int, tau []float64, c []float64, ldc int) error {
	// Query for workspace size.
	work := make([]float64, 1)
	err := dormqrHelper(side, trans, m, n, k, a, lda, tau, c, ldc, work, -1)
	if err != nil {
		return err
	}

	lwork := int(work[0])
	work = make([]float64, lwork)
	return dormqrHelper(side, trans, m, n, k, a, lda, tau, c, ldc, work, lwork)
}

func dormqrHelper(side matSide, trans bool, m, n, k int, a []float64, lda int, tau []float64, c []float64, ldc int, work []float64, lwork int) error {
	var (
		side_  = sideChar(side)
		trans_ = transChar(trans)
		m_     = C.integer(m)
		n_     = C.integer(n)
		k_     = C.integer(k)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		tau_   = ptrFloat64(tau)
		c_     = ptrFloat64(c)
		ldc_   = C.integer(ldc)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dormqr_(&side_, &trans_, &m_, &n_, &k_, a_, &lda_, tau_, c_, &ldc_, work_, &lwork_, &info_)

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
