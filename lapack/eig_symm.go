package lapack

// Computes the eigenvalue factorization of a symmetric matrix.
// Calls DSYEV.
//
// There is no function to compute the eigenvectors of a general matrix
// in this library because matrices with real eigenvectors are symmetric.
func EigSymm(a Const) (*Mat, []float64, error) {
	if err := errNonPosDims(a); err != nil {
		return nil, nil, err
	}
	if err := errNonSquare(a); err != nil {
		return nil, nil, err
	}
	if err := errNonSymm(a); err != nil {
		return nil, nil, err
	}
	return eigSymm(cloneMat(a), DefaultTri)
}

func eigSymm(a *Mat, tri Triangle) (*Mat, []float64, error) {
	n, _ := a.Dims()
	d, err := dsyev(vectors, UpperTri, n, a.Elems, n)
	if err != nil {
		return nil, nil, err
	}
	return a, d, nil
}
