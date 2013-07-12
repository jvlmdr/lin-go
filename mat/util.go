package mat

func sort(a, b int) (int, int) {
	if b < a {
		return b, a
	}
	return a, b
}

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
