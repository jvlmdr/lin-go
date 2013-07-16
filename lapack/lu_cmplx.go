package lapack

import "github.com/jackvalmadre/lin-go/zmat"

// #include "../f2c.h"
import "C"

// Describes a real LU factorization.
type ComplexLU struct {
	A    zmat.ColMajor
	Ipiv []C.integer
}
