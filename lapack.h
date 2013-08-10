/* http://www.netlib.org/clapack/clapack.h */

// Real

/* Subroutine */ int dgels_(char *trans, integer *m, integer *n, integer *
	nrhs, doublereal *a, integer *lda, doublereal *b, integer *ldb, 
	doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dgelsd_(integer *m, integer *n, integer *nrhs, 
	doublereal *a, integer *lda, doublereal *b, integer *ldb, doublereal *
	s, doublereal *rcond, integer *rank, doublereal *work, integer *lwork, 
	 integer *iwork, integer *info);

/* Subroutine */ int dgesv_(integer *n, integer *nrhs, doublereal *a, integer 
	*lda, integer *ipiv, doublereal *b, integer *ldb, integer *info);

/* Subroutine */ int dsysv_(char *uplo, integer *n, integer *nrhs, doublereal 
	*a, integer *lda, integer *ipiv, doublereal *b, integer *ldb, 
	doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dsytrf_(char *uplo, integer *n, doublereal *a, integer *
	lda, integer *ipiv, doublereal *work, integer *lwork, integer *info);

/* Subroutine */ int dsytrs_(char *uplo, integer *n, integer *nrhs, 
	doublereal *a, integer *lda, integer *ipiv, doublereal *b, integer *
	ldb, integer *info);

// Complex

/* Subroutine */ int zgels_(char *trans, integer *m, integer *n, integer *
	nrhs, doublecomplex *a, integer *lda, doublecomplex *b, integer *ldb, 
	doublecomplex *work, integer *lwork, integer *info);

/* Subroutine */ int zgelsd_(integer *m, integer *n, integer *nrhs, 
	doublecomplex *a, integer *lda, doublecomplex *b, integer *ldb, 
	doublereal *s, doublereal *rcond, integer *rank, doublecomplex *work, 
	integer *lwork, doublereal *rwork, integer *iwork, integer *info);

/* Subroutine */ int zgesv_(integer *n, integer *nrhs, doublecomplex *a, 
	integer *lda, integer *ipiv, doublecomplex *b, integer *ldb, integer *
	info);
