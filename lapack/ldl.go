package lapack

// Describes an LDL factorization (Cholesky with a diagonal matrix).
type LDLFact struct {
	A   *Mat
	Tri Triangle
	Piv []int
}

// Computes an LDL factorization.
// Calls DSYTRF.
// Equivalent to SolveSymm (calls DSYSV).
func LDL(a Const) (*LDLFact, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errNonSymm(a); err != nil {
		return nil, err
	}
	return ldl(cloneMat(a), DefaultTri)
}

func ldl(a *Mat, tri Triangle) (*LDLFact, error) {
	n, _ := a.Dims()
	piv, err := dsytrf(tri, n, a.Elems, n)
	if err != nil {
		return nil, err
	}
	return &LDLFact{a, tri, piv}, nil
}

// Solves a square, symmetric system given its LDL factorization.
// Calls DSYTRS.
func (ldl *LDLFact) Solve(b []float64) ([]float64, error) {
	if err := errIncompat(ldl.A, b); err != nil {
		return nil, err
	}
	return ldl.solve(cloneSlice(b))
}

// b will be modified.
func (ldl *LDLFact) solve(b []float64) ([]float64, error) {
	n, _ := ldl.A.Dims()
	err := dsytrs(ldl.Tri, n, 1, ldl.A.Elems, n, ldl.Piv, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
