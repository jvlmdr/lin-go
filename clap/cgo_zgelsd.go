package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZGELSD: complex double-precision GEneral Least Squares (svd, Divide-and-conquer)
//
// http://www.netlib.org/lapack/complex16/zgelsd.f
func zgelsd(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, rcond float64) error {
	// Singular values.
	s := make([]float64, min(m, n))

	// Request workspace size.
	var (
		work  = make([]complex128, 1)
		rwork = make([]float64, 1)
		iwork = make([]C.integer, 1)
	)
	err := zgelsdHelper(m, n, nrhs, a, lda, b, ldb, s, rcond, work, -1, rwork, iwork)
	if err != nil {
		return err
	}

	lwork := int(real(work[0]))
	lrwork := int(rwork[0])
	liwork := int(iwork[0])
	work = make([]complex128, max(1, lwork))
	rwork = make([]float64, max(1, lrwork))
	iwork = make([]C.integer, max(1, liwork))
	return zgelsdHelper(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, rwork, iwork)
}

func zgelsdHelper(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64, work []complex128, lwork int, rwork []float64, iwork []C.integer) error {
	var (
		m_     = C.integer(m)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		b_     = ptrComplex128(b)
		ldb_   = C.integer(ldb)
		s_     = ptrFloat64(s)
		rcond_ = C.doublereal(rcond)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
		rwork_ = ptrFloat64(rwork)
		iwork_ = ptrInt(iwork)
	)
	var (
		rank_ C.integer
		info_ C.integer
	)

	C.zgelsd_(&m_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, s_, &rcond_, &rank_, work_, &lwork_, rwork_, iwork_, &info_)

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
