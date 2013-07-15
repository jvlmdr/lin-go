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

func forceToReal(x interface{}) float64 {
	if y, ok := x.(float64); ok {
		return y
	}
	if y, ok := x.(complex128); ok {
		return real(y)
	}
	panic("Neither float64 not complex128")
}
