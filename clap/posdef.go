package clap

// Solves A x = b where A is Hermitian and positive-definite.
// Calls ZPOSV.
func SolvePosDef(a Const, b []complex128) ([]complex128, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errIncompat(a, b); err != nil {
		return nil, err
	}
	return solvePosDef(cloneMat(a), cloneSlice(b))
}

// a and b will be modified.
func solvePosDef(a *Mat, b []complex128) ([]complex128, error) {
	n, _ := a.Dims()
	err := zposv(n, 1, a.Elems, n, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}
