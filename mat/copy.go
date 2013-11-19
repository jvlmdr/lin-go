package mat

func Copy(dst Mutable, src Const) {
	if err := errIfDimsNotEq(src, dst); err != nil {
		panic(err)
	}

	m, n := src.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dst.Set(i, j, src.At(i, j))
		}
	}
}
