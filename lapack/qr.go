package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"

	"fmt"
)

type QRFact struct {
	A   mat.Stride
	Tau []float64
}

func QR(A mat.Const) (QRFact, error) {
	return QRNoCopy(mat.MakeStrideCopy(A))
}

func QRNoCopy(A mat.Stride) (QRFact, error) {
	tau := make([]float64, min(A.Rows, A.Cols))
	info := dgeqrfAuto(A.Rows, A.Cols, A.Elems, A.Stride, tau)
	if info != 0 {
		return QRFact{}, ErrNonZeroInfo(info)
	}
	return QRFact{A, tau}, nil
}

// Automatically allocates workspace.
func dgeqrfAuto(m, n int, a []float64, lda int, tau []float64) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info := dgeqrf(m, n, a, lda, tau, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dgeqrf(m, n, a, lda, tau, work, lwork)
}

////////////////////////////////////////////////////////////////////////////////

func (qr QRFact) Solve(T bool, b vec.Const) (vec.Slice, error) {
	return qr.SolveNoCopy(T, vec.MakeSliceCopy(b))
}

func (qr QRFact) SolveMat(T bool, B mat.Const) (mat.Stride, error) {
	return qr.SolveMatNoCopy(T, mat.MakeStrideCopy(B))
}

func (qr QRFact) SolveNoCopy(T bool, b vec.Slice) (vec.Slice, error) {
	B := mat.StrideMat(b)
	X, err := qr.SolveMatNoCopy(T, B)
	if err != nil {
		return vec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B or A^T X = B.
func (qr QRFact) SolveMatNoCopy(T bool, B mat.Stride) (mat.Stride, error) {
	size := qr.A.Size()
	if T {
		size = size.T()
	}

	if size.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", size, B.Size()))
	}

	if !T {
		// Over-constrained system.
		if size.Rows <= size.Cols {
			panic(fmt.Sprintf("system not over-constrained: %v", size))
		}

		var info int
		// B = Q' * B
		info = dormqrAuto(Left, Trans, size.Rows, B.Cols, size.Cols,
			qr.A.Elems, qr.A.Stride, qr.Tau, B.Elems, B.Stride)
		if info != 0 {
			return mat.Stride{}, ErrNonZeroInfo(info)
		}
		// B = R \ B
		info = dtrtrs(UpperTriangle, NoTrans, NonUnitTri, size.Cols, B.Cols,
			qr.A.Elems, qr.A.Stride, B.Elems, B.Stride)
		if info != 0 {
			return mat.Stride{}, ErrNonZeroInfo(info)
		}
		// Take sub-slice of B.
		B = B.SliceTo(size.Cols, B.Cols)
		return B, nil
	} else {
		// Under-constrained system.
		if size.Cols <= size.Rows {
			panic(fmt.Sprintf("system not under-constrained: %v", size))
		}

		var info int
		// B = R' \ B
		info = dtrtrs(UpperTriangle, Trans, NonUnitTri, size.Rows, B.Cols,
			qr.A.Elems, qr.A.Stride, B.Elems, B.Stride)
		if info != 0 {
			return mat.Stride{}, ErrNonZeroInfo(info)
		}
		// Append rows of zero to B.
		B = appendRows(B, size.Cols-B.Rows)
		// B = Q * B
		info = dormqrAuto(Left, NoTrans, size.Cols, B.Cols, size.Rows,
			qr.A.Elems, qr.A.Stride, qr.Tau, B.Elems, B.Stride)
		if info != 0 {
			return mat.Stride{}, ErrNonZeroInfo(info)
		}
	}

	return B, nil
}

// Automatically allocates workspace.
func dormqrAuto(side Side, trans Transpose, m, n, k int, a []float64, lda int, tau []float64, c []float64, ldc int) int {
	var (
		lwork = -1
		work  = make([]float64, 1)
	)
	info := dormqr(side, trans, m, n, k, a, lda, tau, c, ldc, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(work[0])
	work = nil
	if lwork > 0 {
		work = make([]float64, lwork)
	}
	return dormqr(side, trans, m, n, k, a, lda, tau, c, ldc, work, lwork)
}
