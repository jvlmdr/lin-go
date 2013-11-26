package clap

// QR factorization.
type QRFact struct {
	A   *Mat
	Tau []complex128
}

// Computes QR factorization.
// Calls ZGEQRF.
func QR(a Const) (*QRFact, error) {
	return qr(cloneMat(a))
}

// a will be modified.
func qr(a *Mat) (*QRFact, error) {
	m, n := a.Dims()
	tau := make([]complex128, min(m, n))

	err := zgeqrf(m, n, a.Elems, m, tau)
	if err != nil {
		return nil, err
	}
	return &QRFact{a, tau}, nil
}

// Solves a linear system using QR decomposition.
// Matrix must be m x n with m >= n (i.e. "skinny")
// for R to be full rank and square.
//
// If h is false, finds x which minimizes ||A x - b|| = ||Q R x - b||.
// Computes R^-1 (Q' b).
//
// If h is true, finds minimum norm x which satisfies b = A' x = R' Q' x.
// Computes Q (R^-T b).
func (qr *QRFact) Solve(h bool, b []complex128) ([]complex128, error) {
	if err := errIncompatT(qr.A, h, b); err != nil {
		return nil, err
	}
	m, n := qr.A.Dims()
	return qr.solve(h, cloneSliceCap(b, max(m, n)))
}

// b will be modified.
// b must have capacity for solution.
func (qr *QRFact) solve(h bool, b []complex128) ([]complex128, error) {
	m, n := qr.A.Dims()
	// In order to be able to solve systems with a QR factorization,
	// the matrix must be skinny (otherwise use LQ, or QR of transpose).
	if m < n {
		return nil, errBadShape(m, n)
	}

	if !h {
		// Q R x = b
		// x = R \ (Q' b)
		var err error

		// b <- Q' b
		err = zunmqr(left, true, m, 1, n, qr.A.Elems, m, qr.Tau, b, len(b))
		if err != nil {
			return nil, err
		}
		// b <- R \ b
		err = ztrtrs(UpperTri, false, nonUnitDiag, n, 1, qr.A.Elems, m, b, len(b))
		if err != nil {
			return nil, err
		}
		// Shrink b to size of solution if necessary.
		b = b[:n]
	} else {
		// R' Q' x = b
		// x = Q (R' \ b)
		var err error

		// Grow b to size of solution if necessary.
		b = b[:m]
		// b <- R' \ b
		err = ztrtrs(UpperTri, true, nonUnitDiag, n, 1, qr.A.Elems, m, b, len(b))
		if err != nil {
			return nil, err
		}
		// b <- Q b
		err = zunmqr(left, false, m, 1, n, qr.A.Elems, m, qr.Tau, b, len(b))
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}
