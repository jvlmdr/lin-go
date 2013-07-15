package lapack

import "C"

type TransposeMode C.char

const (
	NoTrans   = TransposeMode(C.char('N'))
	Trans     = TransposeMode(C.char('T'))
	ConjTrans = TransposeMode(C.char('C'))
)
