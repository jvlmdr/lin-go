package lapack

import "github.com/jackvalmadre/lin-go/mat"

// #include "../f2c.h"
import "C"

// Describes a real LU factorization.
type RealLU struct {
	A    mat.SemiContiguousColMajor
	Ipiv []C.integer
}
