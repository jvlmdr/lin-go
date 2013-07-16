package blas

import "C"

type Transpose C.char

const (
	NoTrans   = Transpose(C.char('N'))
	Trans     = Transpose(C.char('T'))
	ConjTrans = Transpose(C.char('C'))
)
