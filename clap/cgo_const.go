package clap

import "fmt"

import "C"

func conjTransChar(h bool) C.char {
	if h {
		return C.char('C')
	}
	return C.char('N')
}

func uploChar(tri Triangle) C.char {
	switch tri {
	case UpperTri:
		return C.char('U')
	case LowerTri:
		return C.char('L')
	default:
		panic(fmt.Sprintf("invalid uplo value: %v", rune(tri)))
	}
}

func sideChar(side matSide) C.char {
	switch side {
	case left:
		return C.char('L')
	case right:
		return C.char('R')
	default:
		panic(fmt.Sprintf("invalid side value: %v", rune(side)))
	}
}

func diagChar(diag diagType) C.char {
	switch diag {
	case nonUnitDiag:
		return C.char('N')
	case unitDiag:
		return C.char('U')
	default:
		panic(fmt.Sprintf("invalid diag value: %v", rune(diag)))
	}
}

func jobzChar(jobz jobzMode) C.char {
	switch jobz {
	case values:
		return C.char('N')
	case vectors:
		return C.char('V')
	default:
		panic(fmt.Sprintf("invalid jobz value: %v", rune(jobz)))
	}
}
