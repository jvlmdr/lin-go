/*
Basic bindings for doing linear algebra with LAPACK.

Uses the following routines to solve linear systems:
	dgesv     LU          square matrix
	dgels     QR          full rank matrix
	dgelsd    SVD         general matrix
	dposv     Cholesky    square, symmetric, positive-definite matrix
	dsysv     LDL         square, symmetric matrix
and provides access to the following routines for computing and using decompositions respectively:
	dgetrf/s    LU
	dpotrf/s    Cholesky
	dsytrf/s    LDL

No support for banded or triangular matrices.
No support for packed representations.
*/
package lapack
