/*
Basic bindings for doing linear algebra with LAPACK.

Uses the following routines to solve linear systems:
	SolveFullRank    dgels     QR/LQ    full-rank matrix (min norm or min residual)
	SolveCond        dgelsd    SVD      general matrix (threshold singular values)
	SolveSquare      dgesv     LU       full-rank, square matrix
	SolveSymm        dsysv     LDL      full-rank, square, symmetric matrix
	SolvePosDef      dposv     Chol     full-rank, square, symmetric, positive-definite matrix
and provides access to the following routines for computing and using decompositions:
	dgetr[fs]      LU
	dge(qr|lq)f    QR/LQ
	dgesdd         SVD
	dpotr[fs]      Chol
	dsytr[fs]      LDL

No support for banded or triangular matrices.
No support for packed representations.
*/
package lapack
