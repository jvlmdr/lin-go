package lapack

// Solves A x = B where A is full rank.
// Whether the problem is minimum residual or minimum norm
// then depends on the dimension of A.
// Calls DGELS.
func SolveFullRank(a Const, b []float64) ([]float64, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errIncompat(a, b); err != nil {
		return nil, err
	}
	m, n := a.Dims()
	return solveFullRank(cloneMat(a), cloneSliceCap(b, max(m, n)))
}

// a and b will be modified.
// b must have capacity for solution.
func solveFullRank(a *Mat, b []float64) ([]float64, error) {
	m, n := a.Dims()
	b = b[:max(m, n)]
	err := dgels(m, n, 1, a.Elems, m, b, len(b))
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
