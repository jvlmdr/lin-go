package lapack

// Computes thin SVD of an m x n matrix.
// Only the first min(m, n) vectors are computed.
func SVD(a Const) (u *Mat, s []float64, vt *Mat, err error) {
	if err := errNonPosDims(a); err != nil {
		return nil, nil, nil, err
	}
	return svd(cloneMat(a))
}

func svd(a *Mat) (u *Mat, s []float64, vt *Mat, err error) {
	m, n := a.Dims()
	k := min(m, n)
	u = NewMat(m, k)
	s = make([]float64, k)
	vt = NewMat(n, k)

	err = dgesdd(m, n, a.Elems, m, s, u.Elems, m, vt.Elems, n)
	if err != nil {
		return nil, nil, nil, err
	}
	return u, s, vt, nil
}
