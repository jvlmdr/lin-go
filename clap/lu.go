package clap

// Describes an LU factorization with pivoting.
type LUFact struct {
	A   *Mat
	Piv []int
}

// Computes an LU factorization.
// Calls ZGETRF.
// Equivalent to SolveSquare (calls ZGESV).
//
// Note that A does not need to be square to compute the decomposition.
// However, it does need to be square to call Solve().
func LU(a Const) (*LUFact, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	return lu(cloneMat(a))
}

// a will be modified.
func lu(a *Mat) (*LUFact, error) {
	m, n := a.Dims()
	piv, err := zgetrf(m, n, a.Elems, m)
	if err != nil {
		return nil, err
	}
	return &LUFact{a, piv}, nil
}

// Solves A x = b (or A' x = b) where A is square and full-rank
// given its LU factorization.
func (lu *LUFact) Solve(h bool, b []complex128) ([]complex128, error) {
	if err := errNonSquare(lu.A); err != nil {
		return nil, err
	}
	if err := errIncompatT(lu.A, h, b); err != nil {
		return nil, err
	}
	return lu.solve(h, cloneSlice(b))
}

// b will be modified.
func (lu *LUFact) solve(h bool, b []complex128) ([]complex128, error) {
	n, _ := lu.A.Dims()
	err := zgetrs(h, n, 1, lu.A.Elems, n, lu.Piv, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
