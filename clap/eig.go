package clap

// Computes the eigenvalue factorization of a Hermitian matrix.
// Calls ZHEEV.
func EigHerm(a Const) (*Mat, []float64, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, nil, err
	}
	return eigHerm(cloneMat(a), DefaultTri)
}

func eigHerm(a *Mat, tri Triangle) (*Mat, []float64, error) {
	n, _ := a.Dims()
	d, err := zheev(vectors, UpperTri, n, a.Elems, n)
	if err != nil {
		return nil, nil, err
	}
	return a, d, nil
}

// Computes the eigenvalue factorization of a square matrix.
// Calls ZGEEV.
func Eig(a Const) (*Mat, []complex128, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, nil, err
	}
	return eig(cloneMat(a))
}

func eig(a *Mat) (*Mat, []complex128, error) {
	n, _ := a.Dims()
	v := NewMat(n, n)
	d, err := zgeev(vectors, values, n, a.Elems, n, v.Elems, n, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return a, d, nil
}
