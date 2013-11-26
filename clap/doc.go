/*
Basic bindings for doing complex linear algebra with LAPACK.

Uses the following routines to solve linear systems:
	SolveEps         zgelsd    SVD      general matrix (threshold singular values)
	SolveFullRank    zgels     QR/LQ    full-rank matrix (min norm or min residual)
	SolveSquare      zgesv     LU       full-rank, square matrix
	SolveHerm        zhesv     LDL      full-rank, square, Hermitian matrix
	SolvePosDef      zposv     Chol     full-rank, square, Hermitian, positive-definite matrix
and provides access to the following routines for computing and using decompositions:
	LU      zgetrf zgetrs
	QR      zgeqrf zunmqr ztrtrs
	Chol    zpotrf zpotrs
	LDL     zsytrf zsytrs
	SVD     zgesdd
	Eig     zheev zgeev

No support for banded or triangular matrices.
No support for packed representations.
*/
package clap
