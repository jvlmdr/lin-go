package lapack

import (
	"fmt"
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"
)

// Vector version of SolveCondMat().
func SolveCond(A mat.Const, b vec.Const, rcond float64) (vec.Slice, int, []float64, error) {
	C := mat.MakeStrideCopy(A)
	d := make([]float64, b.Len(), max(b.Len(), mat.Cols(A)))
	vec.Copy(vec.Slice(d), b)
	return SolveCondNoCopy(C, d, rcond)
}

// Like SolveCondMatNoCopy() except A and B are left intact.
func SolveCondMat(A mat.Const, B mat.Const, rcond float64) (mat.Stride, int, []float64, error) {
	m, n := mat.RowsCols(A)
	nrhs := mat.Cols(B)

	C := mat.MakeStrideCopy(A)
	D := mat.MakeStrideCap(m, nrhs, max(m, n), nrhs)
	mat.Copy(D, B)
	return SolveCondMatNoCopy(C, D, rcond)
}

// Vector version of SolveCondMatNoCopy().
func SolveCondNoCopy(A mat.Stride, b vec.Slice, rcond float64) (vec.Slice, int, []float64, error) {
	B := mat.StrideMat(b)
	X, rank, sigma, err := SolveCondMatNoCopy(A, B, rcond)
	if err != nil {
		return vec.Slice{}, 0, nil, err
	}
	return X.Col(0), rank, sigma, nil
}

// Solves A X = B where A is not necessarily full rank.
// The user must specify a relative inverse condition number.
//
// Calls dgelsd.
//
// Returns solution, rank and singular values.
// Solution references same data as B.
func SolveCondMatNoCopy(A mat.Stride, B mat.Stride, rcond float64) (mat.Stride, int, []float64, error) {
	// Check dimensions.
	if A.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", A.Size(), B.Size()))
	}

	// Re-allocate B if more space required.
	if !B.InCapTo(A.Cols, B.Cols) {
		rows := max(2*B.Stride, A.Cols)
		cols := B.Cols
		tmp := mat.MakeStrideCap(B.Rows, cols, rows, cols)
		mat.Copy(tmp, B)
		B = tmp
	}

	sigma := make([]float64, min(A.Rows, A.Cols))

	rank, info := dgelsdAuto(A.Rows, A.Cols, B.Cols, A.Elems, A.Stride, B.Elems, B.Stride, sigma, rcond)
	if info != 0 {
		return mat.Stride{}, 0, nil, ErrNonZeroInfo(info)
	}

	X := B.SliceTo(A.Cols, B.Cols)
	return X, rank, sigma, nil
}

// Automatically allocates workspace.
func dgelsdAuto(m, n, nrhs int, a []float64, lda int, b []float64, ldb int, s []float64, rcond float64) (rank int, info int) {
	var (
		lwork = -1
		work  = make([]float64, 1)
		iwork = make(IntList, 1)
	)
	rank, info = dgelsd(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, iwork)
	if info != 0 {
		return
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}

	liwork := int(iwork[0])
	iwork = nil
	if liwork > 0 {
		iwork = make(IntList, liwork)
	}

	return dgelsd(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, iwork)
}

// Complex version of SolveCond().
func SolveCondCmplx(A zmat.Const, b zvec.Const, rcond float64) (zvec.Slice, int, []float64, error) {
	C := zmat.MakeStrideCopy(A)
	d := make([]complex128, b.Len(), max(b.Len(), zmat.Cols(A)))
	zvec.Copy(zvec.Slice(d), b)
	return SolveCondCmplxNoCopy(C, d, rcond)
}

// Complex version of SolveCondMat().
func SolveCondMatCmplx(A zmat.Const, B zmat.Const, rcond float64) (zmat.Stride, int, []float64, error) {
	m, n := zmat.RowsCols(A)
	nrhs := zmat.Cols(B)

	C := zmat.MakeStrideCopy(A)
	D := zmat.MakeStrideCap(m, nrhs, max(m, n), nrhs)
	zmat.Copy(D, B)
	return SolveCondMatCmplxNoCopy(C, D, rcond)
}

// Complex version of SolveCondNoCopy().
func SolveCondCmplxNoCopy(A zmat.Stride, b zvec.Slice, rcond float64) (zvec.Slice, int, []float64, error) {
	B := zmat.StrideMat(b)
	X, rank, sigma, err := SolveCondMatCmplxNoCopy(A, B, rcond)
	if err != nil {
		return zvec.Slice{}, 0, nil, err
	}
	return X.Col(0), rank, sigma, nil
}

// Complex version of SolveCondMatNoCopy().
//
// Calls zgelsd.
func SolveCondMatCmplxNoCopy(A zmat.Stride, B zmat.Stride, rcond float64) (zmat.Stride, int, []float64, error) {
	size := A.Size()
	// Check that B has enough space to contain input and solution.
	if zmat.Rows(B) < size.Rows {
		panic("Not enough rows to contain constraints")
	}
	if zmat.Rows(B) < size.Cols {
		panic("Not enough rows to contain solution")
	}

	sigma := make([]float64, min(size.Rows, size.Cols))

	rank, info := zgelsdAuto(A.Rows, A.Cols, B.Cols, A.Elems, A.Stride, B.Elems, B.Stride, sigma, rcond)
	if info != 0 {
		return zmat.Stride{}, 0, nil, ErrNonZeroInfo(info)
	}

	X := B.SliceTo(A.Cols, B.Cols)
	return X, rank, sigma, nil
}

// Automatically allocates workspace.
func zgelsdAuto(m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int, s []float64, rcond float64) (rank int, info int) {
	var (
		lwork = -1
		work  = make([]complex128, 1)
		rwork = make([]float64, 1)
		iwork = make(IntList, 1)
	)
	rank, info = zgelsd(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, rwork, iwork)
	if info != 0 {
		return
	}

	lwork = int(real(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}

	lrwork := int(rwork[0])
	rwork = nil
	if lrwork > 0 {
		rwork = make([]float64, lrwork)
	}

	liwork := int(iwork[0])
	iwork = nil
	if liwork > 0 {
		iwork = make(IntList, liwork)
	}

	return zgelsd(m, n, nrhs, a, lda, b, ldb, s, rcond, work, lwork, rwork, iwork)
}
