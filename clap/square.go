package clap

// Solves A x = b where A is square and full-rank.
// Calls ZGESV.
func SolveSquare(a Const, b []complex128) ([]complex128, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errIncompat(a, b); err != nil {
		return nil, err
	}
	return solveSquare(cloneMat(a), cloneSlice(b))
}

// a and b will be modified.
func solveSquare(a *Mat, b []complex128) ([]complex128, error) {
	n, _ := a.Dims()
	err := zgesv(n, 1, a.Elems, n, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
