/*
Package lapack links to native library for real linear algebra.

Uses the following routines to solve linear systems:
	SolveEps         dgelsd    SVD      general matrix (threshold singular values)
	SolveFullRank    dgels     QR/LQ    full-rank matrix (min norm or min residual)
	SolveSquare      dgesv     LU       full-rank, square matrix
	SolveSymm        dsysv     LDL      full-rank, square, symmetric matrix
	SolvePosDef      dposv     Chol     full-rank, square, symmetric, positive-definite matrix
and provides access to the following routines for computing and using decompositions:
	LU      dgetrf dgetrs
	QR      dgeqrf dormqr dtrtrs
	Chol    dpotrf dpotrs
	LDL     dsytrf dsytrs
	SVD     dgesdd
	Eig     dsyev

No support for banded or triangular matrices.
No support for packed representations.
*/
package lapack
