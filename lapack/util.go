package lapack

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func re(x interface{}) float64 {
	switch x := x.(type) {
	default:
	case float64:
		return x
	case complex128:
		return real(x)
	}
	panic("neither float64 not complex128")
}
