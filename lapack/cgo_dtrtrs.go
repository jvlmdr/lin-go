package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DTRTRS: (Double-precision) TRiangular Solve
//
// http://www.netlib.org/lapack/double/dtrtrs.f
func dtrtrs(tri Triangle, t bool, diag diagType, n, nrhs int, a []float64, lda int, b []float64, ldb int) error {
	var (
		uplo_  = uploChar(tri)
		trans_ = transChar(t)
		diag_  = diagChar(diag)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrFloat64(a)
		lda_   = C.integer(lda)
		b_     = ptrFloat64(b)
		ldb_   = C.integer(ldb)
	)
	var info_ C.integer

	C.dtrtrs_(&uplo_, &trans_, &diag_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)

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
