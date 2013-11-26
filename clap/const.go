package clap

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
