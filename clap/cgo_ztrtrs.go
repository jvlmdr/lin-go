package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DTRTRS: complex double-precision TRiangular Solve
//
// http://www.netlib.org/lapack/complex16/ztrtrs.f
func ztrtrs(tri Triangle, h bool, diag diagType, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	var (
		uplo_  = uploChar(tri)
		trans_ = conjTransChar(h)
		diag_  = diagChar(diag)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		b_     = ptrComplex128(b)
		ldb_   = C.integer(ldb)
	)
	var info_ C.integer

	C.ztrtrs_(&uplo_, &trans_, &diag_, &n_, &nrhs_, a_, &lda_, b_, &ldb_, &info_)

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
