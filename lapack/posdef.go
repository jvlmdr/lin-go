package lapack

// Solves A x = b where A is symmetric and positive-definite.
// Calls DPOSV.
func SolvePosDef(a Const, b []float64) ([]float64, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errIncompat(a, b); err != nil {
		return nil, err
	}
	if err := errNonSymm(a); err != nil {
		return nil, err
	}
	return solvePosDef(cloneMat(a), cloneSlice(b))
}

// a and b will be modified.
func solvePosDef(a *Mat, b []float64) ([]float64, error) {
	n, _ := a.Dims()
	err := dposv(n, 1, a.Elems, n, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
