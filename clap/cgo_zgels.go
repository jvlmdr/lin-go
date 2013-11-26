package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGELS: complex double-precision GEneral Least Squares
//
// http://www.netlib.org/lapack/complex16/zgels.f
func zgels(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	// Query workspace size.
	work := make([]complex128, 1)
	err := zgelsHelper(m, n, nrhs, a, lda, b, ldb, work, -1)
	if err != nil {
		return err
	}

	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	return zgelsHelper(m, n, nrhs, a, lda, b, ldb, work, lwork)
}

func zgelsHelper(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, work []complex128, lwork int) error {
	var (
		trans_ = conjTransChar(false)
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		b_     = ptrComplex128(b)
		ldb_   = C.integer(ldb)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.zgels_(&trans_, &m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, work_, &lwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errNotFullRank(info)
	default:
		return nil
	}
}
