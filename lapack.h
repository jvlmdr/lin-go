/* http://www.netlib.org/clapack/clapack.h */

// Real

/* Subroutine */ int dgels_(char *trans, integer *m, integer *n, integer *
	nrhs, doublereal *a, integer *lda, doublereal *b, integer *ldb,
	doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dgelsd_(integer *m, integer *n, integer *nrhs,
	doublereal *a, integer *lda, doublereal *b, integer *ldb, doublereal *
	s, doublereal *rcond, integer *rank, doublereal *work, integer *lwork,
	 integer *iwork, integer *info);

/* Subroutine */ int dgeqrf_(integer *m, integer *n, doublereal *a, integer *
	lda, doublereal *tau, doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dgesv_(integer *n, integer *nrhs, doublereal *a, integer
	*lda, integer *ipiv, doublereal *b, integer *ldb, integer *info);

/* Subroutine */ int dgetrf_(integer *m, integer *n, doublereal *a, integer *
	lda, integer *ipiv, integer *info);

/* Subroutine */ int dgetri_(integer *n, doublereal *a, integer *lda, integer
	*ipiv, doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dgetrs_(char *trans, integer *n, integer *nrhs,
	doublereal *a, integer *lda, integer *ipiv, doublereal *b, integer *
	ldb, integer *info);

/* Subroutine */ int dormqr_(char *side, char *trans, integer *m, integer *n,
	integer *k, doublereal *a, integer *lda, doublereal *tau, doublereal *
	c__, integer *ldc, doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dposv_(char *uplo, integer *n, integer *nrhs, doublereal
	*a, integer *lda, doublereal *b, integer *ldb, integer *info);

/* Subroutine */ int dpotrf_(char *uplo, integer *n, doublereal *a, integer *
	lda, integer *info);

/* Subroutine */ int dpotrs_(char *uplo, integer *n, integer *nrhs,
	doublereal *a, integer *lda, doublereal *b, integer *ldb, integer *
	info);

/* Subroutine */ int dsyev_(char *jobz, char *uplo, integer *n, doublereal *a,
	integer *lda, doublereal *w, doublereal *work, integer *lwork,
	integer *info);

/* Subroutine */ int dsysv_(char *uplo, integer *n, integer *nrhs, doublereal
	*a, integer *lda, integer *ipiv, doublereal *b, integer *ldb,
	doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dsytrf_(char *uplo, integer *n, doublereal *a, integer *
	lda, integer *ipiv, doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dsytrs_(char *uplo, integer *n, integer *nrhs,
	doublereal *a, integer *lda, integer *ipiv, doublereal *b, integer *
	ldb, integer *info);

/* Subroutine */ int dtrtrs_(char *uplo, char *trans, char *diag, integer *n,
	integer *nrhs, doublereal *a, integer *lda, doublereal *b, integer *
	ldb, integer *info);

// Complex

/* Subroutine */ int zgels_(char *trans, integer *m, integer *n, integer *
	nrhs, doublecomplex *a, integer *lda, doublecomplex *b, integer *ldb,
	doublecomplex *work, integer *lwork, integer *info);

/* Subroutine */ int zgelsd_(integer *m, integer *n, integer *nrhs,
	doublecomplex *a, integer *lda, doublecomplex *b, integer *ldb,
	doublereal *s, doublereal *rcond, integer *rank, doublecomplex *work,
	integer *lwork, doublereal *rwork, integer *iwork, integer *info);

/* Subroutine */ int zgeqrf_(integer *m, integer *n, doublecomplex *a,
	integer *lda, doublecomplex *tau, doublecomplex *work, integer *lwork,
	integer *info);

/* Subroutine */ int zgesv_(integer *n, integer *nrhs, doublecomplex *a,
	integer *lda, integer *ipiv, doublecomplex *b, integer *ldb, integer *
	info);

/* Subroutine */ int zgetrf_(integer *m, integer *n, doublecomplex *a,
	integer *lda, integer *ipiv, integer *info);

/* Subroutine */ int zgetri_(integer *n, doublecomplex *a, integer *lda,
	integer *ipiv, doublecomplex *work, integer *lwork, integer *info);

/* Subroutine */ int zgetrs_(char *trans, integer *n, integer *nrhs,
	doublecomplex *a, integer *lda, integer *ipiv, doublecomplex *b,
	integer *ldb, integer *info);

/* Subroutine */ int zsysv_(char *uplo, integer *n, integer *nrhs,
	doublecomplex *a, integer *lda, integer *ipiv, doublecomplex *b,
	integer *ldb, doublecomplex *work, integer *lwork, integer *info);

/* Subroutine */ int zsytrf_(char *uplo, integer *n, doublecomplex *a,
	integer *lda, integer *ipiv, doublecomplex *work, integer *lwork,
	integer *info);

/* Subroutine */ int zsytri_(char *uplo, integer *n, doublecomplex *a,
	integer *lda, integer *ipiv, doublecomplex *work, integer *info);

/* Subroutine */ int zsytrs_(char *uplo, integer *n, integer *nrhs,
	doublecomplex *a, integer *lda, integer *ipiv, doublecomplex *b,
	integer *ldb, integer *info);

/* Subroutine */ int ztrtrs_(char *uplo, char *trans, char *diag, integer *n,
	integer *nrhs, doublecomplex *a, integer *lda, doublecomplex *b,
	integer *ldb, integer *info);

/* Subroutine */ int zunmqr_(char *side, char *trans, integer *m, integer *n,
	integer *k, doublecomplex *a, integer *lda, doublecomplex *tau,
	doublecomplex *c__, integer *ldc, doublecomplex *work, integer *lwork,
	integer *info);
