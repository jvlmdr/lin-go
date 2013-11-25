package lapack

import "runtime"

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DSYSV: (Double-precision) SYmmetric SolVe
//
// http://www.netlib.org/lapack/double/dsysv.f
func dsysv(n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	ipiv := make([]C.integer, n)

	// Request workspace size.
	work := make([]float64, 1)
	err := dsysvHelper(n, nrhs, a, lda, ipiv, b, ldb, work, -1)
	if err != nil {
		return err
	}

	// Allocate workspace and make call.
	lwork := int(work[0])
	work = make([]float64, lwork)
	return dsysvHelper(n, nrhs, a, lda, ipiv, b, ldb, work, lwork)
}

// Needs to be supplied ipiv and work.
func dsysvHelper(n, nrhs int, a []float64, lda int, ipiv []C.integer, b []float64, ldb int, work []float64, lwork int) error {
	defer runtime.GC()

	var (
		uplo_  = C.char(DefaultTri)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		b_     = ptrFloat64(b)
		ldb_   = C.integer(ldb)
		work_  = ptrFloat64(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.dsysv_(&uplo_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, work_, &lwork_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errSingular(info)
	default:
		return nil
	}
}
