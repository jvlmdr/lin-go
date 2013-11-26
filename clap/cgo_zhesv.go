package clap

// #include "f2c.h"
// #include "clapack.h"
import "C"

// ZHESV: complex double-precision HErmitian SolVe
//
// http://www.netlib.org/lapack/complex16/zhesv.f
func zhesv(n, nrhs int, a []complex128, lda int, b []complex128, ldb int) error {
	ipiv := make([]C.integer, n)

	// Request workspace size.
	work := make([]complex128, 1)
	err := zhesvHelper(n, nrhs, a, lda, ipiv, b, ldb, work, -1)
	if err != nil {
		return err
	}

	// Allocate workspace and make call.
	lwork := int(real(work[0]))
	work = make([]complex128, max(1, lwork))
	return zhesvHelper(n, nrhs, a, lda, ipiv, b, ldb, work, lwork)
}

// Needs to be supplied ipiv and work.
func zhesvHelper(n, nrhs int, a []complex128, lda int, ipiv []C.integer, b []complex128, ldb int, work []complex128, lwork int) error {
	var (
		uplo_  = C.char(DefaultTri)
		n_     = C.integer(n)
		nrhs_  = C.integer(nrhs)
		a_     = ptrComplex128(a)
		lda_   = C.integer(lda)
		ipiv_  = ptrInt(ipiv)
		b_     = ptrComplex128(b)
		ldb_   = C.integer(ldb)
		work_  = ptrComplex128(work)
		lwork_ = C.integer(lwork)
	)
	var info_ C.integer

	C.zhesv_(&uplo_, &n_, &nrhs_, a_, &lda_, ipiv_, b_, &ldb_, work_, &lwork_, &info_)

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
