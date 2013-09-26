package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"github.com/jackvalmadre/lin-go/vec"
	"github.com/jackvalmadre/lin-go/zmat"
	"github.com/jackvalmadre/lin-go/zvec"

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

// Solves a linear system using QR decomposition.
// A should be mxn with m > n.
//
// If T is false, finds x which minimizes ||A x - b|| = ||Q R x - b||.
// Computes R^-1 (Q' b).
//
// If T is true, finds minimum norm x which satisfies b = A' x = R' Q' x.
// Computes Q (R^-T b).
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

////////////////////////////////////////////////////////////////////////////////

type QRFactCmplx struct {
	A   zmat.Stride
	Tau []complex128
}

func QRCmplx(A zmat.Const) (QRFactCmplx, error) {
	return QRCmplxNoCopy(zmat.MakeStrideCopy(A))
}

func QRCmplxNoCopy(A zmat.Stride) (QRFactCmplx, error) {
	tau := make([]complex128, min(A.Rows, A.Cols))
	info := zgeqrfAuto(A.Rows, A.Cols, A.Elems, A.Stride, tau)
	if info != 0 {
		return QRFactCmplx{}, ErrNonZeroInfo(info)
	}
	return QRFactCmplx{A, tau}, nil
}

// Automatically allocates workspace.
func zgeqrfAuto(m, n int, a []complex128, lda int, tau []complex128) int {
	var (
		lwork = -1
		work  = make([]complex128, 1)
	)
	info := zgeqrf(m, n, a, lda, tau, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(real(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}
	return zgeqrf(m, n, a, lda, tau, work, lwork)
}

////////////////////////////////////////////////////////////////////////////////

func (qr QRFactCmplx) Solve(H bool, b zvec.Const) (zvec.Slice, error) {
	return qr.SolveNoCopy(H, zvec.MakeSliceCopy(b))
}

func (qr QRFactCmplx) SolveMat(H bool, B zmat.Const) (zmat.Stride, error) {
	return qr.SolveMatNoCopy(H, zmat.MakeStrideCopy(B))
}

func (qr QRFactCmplx) SolveNoCopy(H bool, b zvec.Slice) (zvec.Slice, error) {
	B := zmat.StrideMat(b)
	X, err := qr.SolveMatNoCopy(H, B)
	if err != nil {
		return zvec.Slice{}, err
	}
	return X.Col(0), nil
}

// Solves A X = B or A^H X = B.
func (qr QRFactCmplx) SolveMatNoCopy(H bool, B zmat.Stride) (zmat.Stride, error) {
	size := qr.A.Size()
	if H {
		size = size.T()
	}

	if size.Rows != B.Rows {
		panic(fmt.Sprintf("dimensions incompatible: %v and %v", size, B.Size()))
	}

	if !H {
		// Over-constrained system.
		if size.Rows <= size.Cols {
			panic(fmt.Sprintf("system not over-constrained: %v", size))
		}

		var info int
		// B = Q' * B
		info = zunmqrAuto(Left, ConjTrans, size.Rows, B.Cols, size.Cols,
			qr.A.Elems, qr.A.Stride, qr.Tau, B.Elems, B.Stride)
		if info != 0 {
			return zmat.Stride{}, ErrNonZeroInfo(info)
		}
		// B = R \ B
		info = ztrtrs(UpperTriangle, NoTrans, NonUnitTri, size.Cols, B.Cols,
			qr.A.Elems, qr.A.Stride, B.Elems, B.Stride)
		if info != 0 {
			return zmat.Stride{}, ErrNonZeroInfo(info)
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
		info = ztrtrs(UpperTriangle, ConjTrans, NonUnitTri, size.Rows, B.Cols,
			qr.A.Elems, qr.A.Stride, B.Elems, B.Stride)
		if info != 0 {
			return zmat.Stride{}, ErrNonZeroInfo(info)
		}
		// Append rows of zero to B.
		B = appendRowsCmplx(B, size.Cols-B.Rows)
		// B = Q * B
		info = zunmqrAuto(Left, NoTrans, size.Cols, B.Cols, size.Rows,
			qr.A.Elems, qr.A.Stride, qr.Tau, B.Elems, B.Stride)
		if info != 0 {
			return zmat.Stride{}, ErrNonZeroInfo(info)
		}
	}

	return B, nil
}

// Automatically allocates workspace.
func zunmqrAuto(side Side, trans Transpose, m, n, k int, a []complex128, lda int, tau []complex128, c []complex128, ldc int) int {
	var (
		lwork = -1
		work  = make([]complex128, 1)
	)
	info := zunmqr(side, trans, m, n, k, a, lda, tau, c, ldc, work, lwork)
	if info != 0 {
		return info
	}

	lwork = int(real(work[0]))
	work = nil
	if lwork > 0 {
		work = make([]complex128, lwork)
	}
	return zunmqr(side, trans, m, n, k, a, lda, tau, c, ldc, work, lwork)
}
