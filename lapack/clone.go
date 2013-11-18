package lapack

func cloneMat(src Const) *Mat {
	rows, cols := src.Dims()
	dst := NewMat(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dst.Set(i, j, src.At(i, j))
		}
	}
	return dst
}

func cloneSlice(src []float64) []float64 {
	return cloneSliceCap(src, len(src))
}

func cloneSliceCap(src []float64, n int) []float64 {
	dst := make([]float64, len(src), n)
	copy(dst, src)
	return dst
}
