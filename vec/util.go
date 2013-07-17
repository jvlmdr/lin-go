package vec

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func panicIfNotSameLength(xs ...Const) {
	var n int
	for i, x := range xs {
		if i == 0 {
			n = x.Len()
			continue
		}
		if x.Len() != n {
			if len(xs) == 2 {
				panic(fmt.Sprintf("Vectors had different length (%d, %d)", n, x.Len()))
			}
			panic(fmt.Sprintf("Vector %d had different length (%d) to first vector (%d)", i, x.Len(), n))
		}
	}
}
