package lapack

import "C"

type Transpose C.char

const (
	NoTrans   = Transpose(C.char('N'))
	Trans     = Transpose(C.char('T'))
	ConjTrans = Transpose(C.char('C'))
)

type UpLo C.char

const (
	UpperTriangle = UpLo(C.char('U'))
	LowerTriangle = UpLo(C.char('L'))
)

type Side C.char

const (
	Left  = Side(C.char('L'))
	Right = Side(C.char('R'))
)

type Diag C.char

const (
	NonUnitTri = Diag(C.char('N'))
	UnitTri    = Diag(C.char('U'))
)

type Jobz C.char

const (
	Values  = Jobz(C.char('N'))
	Vectors = Jobz(C.char('V'))
)
