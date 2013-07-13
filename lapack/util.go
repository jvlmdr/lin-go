package lapack

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
