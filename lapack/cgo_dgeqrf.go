package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGEQRF: (Double-precision) GEneral QR Factor
//
// http://www.netlib.org/lapack/double/dgeqrf.f
func dgeqrf(m, n int, a []float64, lda int, tau []float64) error {
	// Query workspace size.
	work := make([]float64, 1)
	err := dgeqrfHelper(m, n, a, lda, tau, work, -1)
	if err != nil {
		return err
	}

	lwork := int(work[0])
	work = make([]float64, max(1, lwork))
	return dgeqrfHelper(m, n, a, lda, tau, work, lwork)
}

func dgeqrfHelper(m, n int, a []float64, lda int, tau, work []float64, lwork int) error {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		tau_   = ptrFloat64(tau)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dgeqrf_(&m_, &n_, a_, &lda_, tau_, work_, &lwork_, &info_)

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
