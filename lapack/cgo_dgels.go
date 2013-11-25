package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGELS: (Double-precision) GEneral Least Squares
//
// http://www.netlib.org/lapack/double/dgels.f
func dgels(m, n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	// Query workspace size.
	work := make([]float64, 1)
	err := dgelsHelper(m, n, nrhs, a, lda, b, ldb, work, -1)
	if err != nil {
		return err
	}

	lwork := int(work[0])
	work = make([]float64, lwork)
	return dgelsHelper(m, n, nrhs, a, lda, b, ldb, work, lwork)
}

func dgelsHelper(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) error {
	defer runtime.GC()

	var (
		trans_ = transChar(false)
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		b_     = ptrFloat64(b)
		ldb_   = C.integer(ldb)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dgels_(&trans_, &m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, work_, &lwork_, &info_)

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
