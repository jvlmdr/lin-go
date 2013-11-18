package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DGELSD: (Double-precision) GEneral Least Squares (svd, Divide-and-conquer)
//
// http://www.netlib.org/lapack/double/dgelsd.f
func dgelsd(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, rcond float64) error {
	// Singular values.
	var s []float64
	if m > 0 && n > 0 {
		s = make([]float64, min(m, n))
	}

	// Request workspace size.
	var (
		work  = make([]float64, 1)
		iwork = make([]C.integer, 1)
	)
	err := dgelsdHelper(m, n, nrhs, a, lda, b, ldb, s, rcond, work, -1, iwork)
	if err != nil {
		return err
	}

	lwork := int(work[0])
	liwork := int(iwork[0])
	work = nil
	iwork = nil

	if lwork > 0 {
		work = make([]float64, lwork)
	}
	if liwork > 0 {
		iwork = make([]C.integer, liwork)
	}
	return dgelsdHelper(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, iwork)
}

func dgelsdHelper(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64, work []float64, lwork int, iwork []C.integer) error {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		b_     = ptrFloat64(b)
		ldb_   = C.integer(ldb)
		s_     = ptrFloat64(s)
		rcond_ = C.doublereal(rcond)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
		iwork_ = ptrInt(iwork)
	)
	var (
		rank_ C.integer
		info_ C.integer
	)

	C.dgelsd_(&m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, s_, &rcond_, &rank_, work_, &lwork_, iwork_, &info_)

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
