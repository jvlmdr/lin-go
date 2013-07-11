package zmat

func sort(a, b int) (int, int) {
	if b < a {
		return b, a
	}
	return a, b
}
