package clap

import "github.com/jvlmdr/lin-go/cmat"

// Describes a Cholesky factorization.
type CholFact struct {
	A   *Mat
	Tri Triangle
}

// Computes the Cholesky factorization A = L L' or A = U' U
// of a Hermitian, positive-definite matrix.
// Calls ZPOTRF.
// Equivalent to SolvePosDef (calls ZPOSV).
func Chol(a Const) (*CholFact, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, err
	}
	if err := errNonHerm(a); err != nil {
		return nil, err
	}
	return chol(cloneMat(a), DefaultTri)
}

// a will be modified.
func chol(a *Mat, tri Triangle) (*CholFact, error) {
	n, _ := a.Dims()
	err := zpotrf(tri, n, a.Elems, n)
	if err != nil {
		return nil, err
	}
	return &CholFact{a, tri}, nil
}

// Solves A x = b where A is Hermitian and positive-definite
// given its Cholesky decomposition.
func (chol *CholFact) Solve(b []complex128) ([]complex128, error) {
	if err := errIncompat(chol.A, b); err != nil {
		return nil, err
	}
	return chol.solve(cloneSlice(b))
}

// b will be modified.
func (chol *CholFact) solve(b []complex128) ([]complex128, error) {
	n, _ := chol.A.Dims()
	err := zpotrs(chol.Tri, n, 1, chol.A.Elems, n, b, n)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Inv returns the matrix inverse.
// Calls ZPOTRI.
func (chol *CholFact) Inv() (*Mat, error) {
	a := cloneMat(chol.A)
	n, _ := a.Dims()
	if err := zpotri(chol.Tri, n, a.Elems, n); err != nil {
		return nil, err
	}
	cmat.Copy(a, &triangleMat{*a, chol.Tri})
	return a, nil
}
