package lapack

// dgels functions which are trivially mapped to zgels functions.

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

	"fmt"
)

// Like SolveFullRankNoCopy() except A and b are intact.
func SolveFullRank(A mat.Const, b vec.Const) (vec.Slice, error) {
	n := mat.Cols(A)
	C := mat.MakeContigCopy(A)
	// Copy b with enough room for x.
	d := make([]float64, b.Len(), max(b.Len(), n))
	vec.Copy(vec.Slice(d), b)

	x, err := SolveFullRankNoCopy(C.Stride(), false, d)
	if err != nil {
		return vec.Slice{}, err
	}
	return x, err
}

// Like SolveFullRankMatNoCopy() except A and B are intact.
func SolveFullRankMat(A mat.Const, B mat.Const) (mat.Stride, error) {
	C := mat.MakeContigCopy(A)
	n := max(mat.Cols(A), mat.Rows(B))
	D := mat.MakeStrideCap(mat.Rows(B), mat.Cols(B), n, mat.Cols(B))
	mat.Copy(D, B)
	return SolveFullRankMatNoCopy(C.Stride(), false, D)
}

// Vector version of SolveFullRankMatNoCopy().
func SolveFullRankNoCopy(A mat.Stride, T bool, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	X, err := SolveFullRankMatNoCopy(A, T, B)
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B where A is full rank.
// Whether the problem is one of minimum residual or minimum norm depends on size of A.
//
// Calls dgels.
//
// If B had sufficient capacity to hold solution, returns a matrix which references elements of B.
// A will be over-written with either the LQ or QR factorization.
func SolveFullRankMatNoCopy(A mat.Stride, T bool, B mat.Stride) (mat.Stride, error) {
	// Tranpose dimensions if specified.
	s := A.Size()
	trans := NoTrans
	if T {
		s = s.T()
		trans = Trans
	}
	// Check dimensions.
	if s.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", s, B.Size()))
	}

	// Re-allocate B if more space required.
	if !B.InCapTo(A.Cols, B.Cols) {
		rows := max(2*B.Stride, A.Cols)
		cols := B.Cols
		tmp := mat.MakeStrideCap(B.Rows, cols, rows, cols)
		mat.Copy(tmp, B)
		B = tmp
	}

	info := dgelsAuto(trans, s.Rows, s.Cols, B.Cols, A.Elems, A.Stride, B.Elems, B.Stride)
	if info != 0 {
		return mat.Stride{}, ErrNonZeroInfo(info)
	}

	X := B.SliceTo(A.Size().Cols, mat.Cols(B))
	return X, nil
}

// Automatically allocates workspace.
func dgelsAuto(trans Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info := dgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
}

// Complex version of SolveFullRank().
func SolveFullRankCmplx(A zmat.Const, b zvec.Const) (zvec.Slice, error) {
	n := zmat.Cols(A)
	C := zmat.MakeContigCopy(A)
	// Copy b with enough room for x.
	d := make([]complex128, b.Len(), max(b.Len(), n))
	zvec.Copy(zvec.Slice(d), b)

	x, err := SolveFullRankCmplxNoCopy(C.Stride(), false, d)
	if err != nil {
		return zvec.Slice{}, err
	}
	return x, err
}

// Complex version of SolveFullRankMat().
func SolveFullRankMatCmplx(A zmat.Const, B zmat.Const) (zmat.Stride, error) {
	C := zmat.MakeContigCopy(A)
	n := max(zmat.Cols(A), zmat.Rows(B))
	D := zmat.MakeStrideCap(zmat.Rows(B), zmat.Cols(B), n, zmat.Cols(B))
	zmat.Copy(D, B)
	return SolveFullRankMatCmplxNoCopy(C.Stride(), false, D)
}

// Complex version of SolveFullRankNoCopy().
func SolveFullRankCmplxNoCopy(A zmat.Stride, H bool, b zvec.Slice) (zvec.Slice, error) {
	B := zmat.StrideMat(b)
	X, err := SolveFullRankMatCmplxNoCopy(A, H, B)
	if err != nil {
		return zvec.Slice{}, err
	}
	return X.Col(0), nil
}

// Complex version of SolveFullRankMatNoCopy().
func SolveFullRankMatCmplxNoCopy(A zmat.Stride, H bool, B zmat.Stride) (zmat.Stride, error) {
	// Tranpose dimensions if specified.
	s := A.Size()
	trans := NoTrans
	if H {
		s = s.T()
		trans = ConjTrans
	}
	// Check dimensions.
	if s.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", s, B.Size()))
	}

	// Re-allocate B if more space required.
	if !B.InCapTo(A.Size().Cols, zmat.Rows(B)) {
		rows := max(2*B.Stride, A.Size().Cols)
		cols := B.Size().Cols
		tmp := zmat.MakeStrideCap(B.Size().Rows, cols, rows, cols)
		zmat.Copy(tmp, B)
		B = tmp
	}

	info := zgelsAuto(trans, s.Rows, s.Cols, B.Cols, A.Elems, A.Stride, B.Elems, B.Stride)
	if info != 0 {
		return zmat.Stride{}, ErrNonZeroInfo(info)
	}

	X := B.SliceTo(A.Size().Cols, zmat.Cols(B))
	return X, nil
}

// Automatically allocates workspace.
func zgelsAuto(trans Transpose, m, n, nrhs int, a []complex128, lda int, b []complex128, ldb int) int {
	var (
		lwork = -1
		work  = make([]complex128, 1)
	)
	info := zgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(real(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}
	return zgels(trans, m, n, nrhs, a, lda, b, ldb, work, lwork)
}
