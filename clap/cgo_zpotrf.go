package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZPOTRF: complex double-precision POsitive-definite TRiangular Factor
//
// http://www.netlib.org/lapack/complex16/zpotrf.f
func zpotrf(uplo Triangle, n int, a []complex128, lda int) error {
	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
	)
	var info_ C.integer

	C.zpotrf_(&uplo_, &n_, a_, &lda_, &info_)
	return zpotrfError(int(info_))
}

func zpotrfError(info int) error {
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errNotPosDef(info)
	default:
		return nil
	}
}
