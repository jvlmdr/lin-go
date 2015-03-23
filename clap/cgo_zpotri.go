package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZPOTRI: complex double-precision POsitive-definite TRiangular factor Invert
//
// http://www.netlib.org/lapack/complex16/zpotri.f
func zpotri(uplo Triangle, n int, a []complex128, lda int) error {
	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		a_    = ptrComplex128(a)
		lda_  = C.integer(lda)
	)
	var info_ C.integer

	C.zpotri_(&uplo_, &n_, a_, &lda_, &info_)

	info := int(info_)
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info == 0:
		return nil
	default:
		return errNotPosDef(info)
	}
}
