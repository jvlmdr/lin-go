package lapack

import "github.com/jackvalmadre/lin-go/zmat"

// Describes a real LU factorization.
type ComplexLU struct {
	A    zmat.ColMajor
	Ipiv IntList
}
