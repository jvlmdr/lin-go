package cmat

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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
