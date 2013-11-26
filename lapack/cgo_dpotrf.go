package lapack

// #include "f2c.h"
// #include "clapack.h"
import "C"

// DPOTRF: (Double-precision) POsitive-definite TRiangular Factor
//
// http://www.netlib.org/lapack/double/dpotrf.f
func dpotrf(uplo Triangle, n int, a []float64, lda int) error {
	var (
		uplo_ = uploChar(uplo)
		n_    = C.integer(n)
		a_    = ptrFloat64(a)
		lda_  = C.integer(lda)
	)
	var info_ C.integer

	C.dpotrf_(&uplo_, &n_, a_, &lda_, &info_)
	return dpotrfError(int(info_))
}

func dpotrfError(info int) error {
	switch {
	case info < 0:
		return errInvalidArg(-info)
	case info > 0:
		return errNotPosDef(info)
	default:
		return nil
	}
}
