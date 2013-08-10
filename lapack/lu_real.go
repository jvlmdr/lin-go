package lapack

import "github.com/jackvalmadre/lin-go/mat"

// Describes a real LU factorization.
type RealLU struct {
	A    mat.ColMajor
	Ipiv IntList
}
