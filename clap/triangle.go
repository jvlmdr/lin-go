package clap

import (
	"fmt"
	"math/cmplx"
)

// triangleMat is a full matrix
// with half of its values ignored.
type triangleMat struct {
	Mat
	Tri Triangle
}

func otherTri(i, j int, tri Triangle) bool {
	switch tri {
	case UpperTri:
		if j < i {
			return true
		}
		return false
	case LowerTri:
		if j < i {
			return true
		}
		return false
	default:
		panic(fmt.Sprintf("unknown triangle identifier: %d", tri))
	}
}

func (a *triangleMat) At(i, j int) complex128 {
	swap := otherTri(i, j, a.Tri)
	if swap {
		i, j = j, i
	}
	x := a.Mat.At(i, j)
	if swap {
		x = cmplx.Conj(x)
	}
	return x
}

func (a *triangleMat) Set(i, j int, x complex128) {
	if otherTri(i, j, a.Tri) {
		i, j = j, i
		x = cmplx.Conj(x)
	}
	a.Set(i, j, x)
}
