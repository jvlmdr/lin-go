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
