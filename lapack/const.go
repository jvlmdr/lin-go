package lapack

import "fmt"

// Specifies the upper or lower triangle of a matrix.
type Triangle rune

const (
	UpperTri = Triangle('U')
	LowerTri = Triangle('L')
)

var DefaultTri = UpperTri

// Specifies whether to perform operation on the left or right.
type matSide rune

const (
	left  = matSide('L')
	right = matSide('R')
)

// Specifies whether diagonal elements are unit or non-unit.
type diagType rune

const (
	unitDiag    = diagType('U')
	nonUnitDiag = diagType('N')
)

// Specifies whether to get eigenvectors or just eigenvalues.
type jobzMode rune

const (
	values  = jobzMode('N')
	vectors = jobzMode('V')
)

func copyToOtherTri(a *Mat, src Triangle) {
	switch src {
	case UpperTri:
		// Iterate over upper triangle.
		for i := 0; i < a.Rows; i++ {
			for j := i; j < a.Cols; j++ {
				// Copy upper to lower.
				a.Set(j, i, a.At(i, j))
			}
		}
	case LowerTri:
		// Iterate over upper triangle.
		for i := 0; i < a.Rows; i++ {
			for j := i; j < a.Cols; j++ {
				// Copy lower to upper.
				a.Set(i, j, a.At(j, i))
			}
		}
	default:
		panic(fmt.Sprintf("unknown triangle: %v", src))
	}
}
