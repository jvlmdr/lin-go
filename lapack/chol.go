package lapack

// CholFact describes a Cholesky factorization.
type CholFact struct {
	A   *Mat
	Tri Triangle
}

// Chol computes the Cholesky factorization A = L L' or A = U' U
// of a symmetric, positive-definite matrix.
// Calls DPOTRF.
// Equivalent to SolvePosDef (calls DPOSV).
func Chol(a Const) (*CholFact, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errNonSymm(a); err != nil {
		return nil, err
	}
	return chol(cloneMat(a), DefaultTri)
}

// a will be modified.
func chol(a *Mat, tri Triangle) (*CholFact, error) {
	n, _ := a.Dims()
	err := dpotrf(tri, n, a.Elems, n)
	if err != nil {
		return nil, err
	}
	return &CholFact{a, tri}, nil
}

// Solve finds x such that A x = b where A is symmetric and positive-definite
// given its Cholesky decomposition.
func (chol *CholFact) Solve(b []float64) ([]float64, error) {
	if err := errIncompat(chol.A, b); err != nil {
		return nil, err
	}
	return chol.solve(cloneSlice(b))
}

// b will be modified.
func (chol *CholFact) solve(b []float64) ([]float64, error) {
	n, _ := chol.A.Dims()
	err := dpotrs(chol.Tri, n, 1, chol.A.Elems, n, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// InvertPosDef computes the inverse of a symmetric positive-definite matrix.
// Calls DPOTRF and DPOTRI.
func InvertPosDef(a Const) (*Mat, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errNonSymm(a); err != nil {
		return nil, err
	}
	x := cloneMat(a)
	if err := invertPosDef(x, DefaultTri); err != nil {
		return nil, err
	}
	return x, nil
}

// a will be modified.
func invertPosDef(a *Mat, tri Triangle) error {
	n, _ := a.Dims()
	if err := dpotrf(tri, n, a.Elems, n); err != nil {
		return err
	}
	if err := dpotri(tri, n, a.Elems, n); err != nil {
		return err
	}
	copyToOtherTri(a, tri)
	return nil
}
