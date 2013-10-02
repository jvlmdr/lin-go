package lapack

import "unsafe"

// #include "../f2c.h"
import "C"

func nonEmptyPtrFloat64(x []float64) *C.doublereal {
	if len(x) == 0 {
		return nil
	}
	return (*C.doublereal)(unsafe.Pointer(&x[0]))
}

func nonEmptyPtrInt(x []int) *C.integer {
	if len(x) == 0 {
		return nil
	}
	return (*C.integer)(unsafe.Pointer(&x[0]))
}

func nonEmptyPtrComplex128(x []complex128) *C.doublecomplex {
	if len(x) == 0 {
		return nil
	}
	return (*C.doublecomplex)(unsafe.Pointer(&x[0]))
}
