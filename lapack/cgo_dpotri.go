package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DPOTRI: (Double-precision) POsitive-definite (TRiangular factor) Inverse
//
// http://www.netlib.org/lapack/double/dpotri.f
func dpotri(uplo Triangle, n int, a []float64, lda int) error {
	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
	)
	var info_ C.integer

	C.dpotri_(&uplo_, &n_, a_, &lda_, &info_)
	return dpotriError(int(info_))
}

func dpotriError(info int) error {
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errNotPosDef(info)
	default:
		return nil
	}
}
