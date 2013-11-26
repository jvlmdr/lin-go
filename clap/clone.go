package clap

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

func cloneSlice(src []complex128) []complex128 {
	return cloneSliceCap(src, len(src))
}

func cloneSliceCap(src []complex128, n int) []complex128 {
	dst := make([]complex128, len(src), n)
	copy(dst, src)
	return dst
}
