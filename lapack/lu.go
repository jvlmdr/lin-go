package lapack

// Describes an LU factorization with pivoting.
type LUFact struct {
	A   *Mat
	Piv []int
}

// Computes an LU factorization.
// Calls DGETRF.
// Equivalent to SolveSquare (calls DGESV).
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
	piv, err := dgetrf(m, n, a.Elems, m)
	if err != nil {
		return nil, err
	}
	return &LUFact{a, piv}, nil
}

// Solves A x = b (or A' x = b) where A is square and full-rank
// given its LU factorization.
func (lu *LUFact) Solve(t bool, b []float64) ([]float64, error) {
	if err := errNonSquare(lu.A); err != nil {
		return nil, err
	}
	if err := errIncompatT(lu.A, t, b); err != nil {
		return nil, err
	}
	return lu.solve(t, cloneSlice(b))
}

// b will be modified.
func (lu *LUFact) solve(t bool, b []float64) ([]float64, error) {
	n, _ := lu.A.Dims()
	err := dgetrs(t, n, 1, lu.A.Elems, n, lu.Piv, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
