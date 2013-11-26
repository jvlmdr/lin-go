package clap

var DefaultEps = 1e-12

// Solves A x = b using SolveEps with DefaultEps.
func Solve(a Const, b []complex128) ([]complex128, error) {
	return SolveEps(a, b, DefaultEps)
}

// Solves A x = b.
// The user must specify epsilon (inverse maximum condition number)
// at which the line is drawn between equality constraints and residuals.
// Calls ZGELSD (eps is the "rcond" parameter).
func SolveEps(a Const, b []complex128, eps float64) ([]complex128, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errIncompat(a, b); err != nil {
		return nil, err
	}
	m, n := a.Dims()
	return solveEps(cloneMat(a), cloneSliceCap(b, max(m, n)), eps)
}

// a and b will be modified.
// b must have capacity for solution.
func solveEps(a *Mat, b []complex128, eps float64) ([]complex128, error) {
	m, n := a.Dims()
	b = b[:max(m, n)]
	err := zgelsd(m, n, 1, a.Elems, m, b, len(b), eps)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
