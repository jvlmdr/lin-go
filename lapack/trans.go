package lapack

import "C"

type TransposeMode C.char

const (
	NoTrans   = TransposeMode(C.char('N'))
	Trans     = TransposeMode(C.char('T'))
	ConjTrans = TransposeMode(C.char('C'))
)

//	// Checks that t is a valid transpose mode for real matrices.
//	func (t TransposeMode) ValidReal() bool {
//		return t == NoTrans || t == Trans
//	}
//	
//	// Checks that t is a valid transpose mode for complex matrices.
//	func (t TransposeMode) ValidComplex() bool {
//		return t == NoTrans || t == Trans || t == ConjTrans
//	}
