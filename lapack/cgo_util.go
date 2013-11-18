package lapack

import "unsafe"

// #include "f2c.h"
import "C"

func ptrInt(x []C.integer) *C.integer {
	if len(x) == 0 {
		return nil
	}
	return &x[0]
}

func ptrFloat64(x []float64) *C.doublereal {
	if len(x) == 0 {
		return nil
	}
	return (*C.doublereal)(unsafe.Pointer(&x[0]))
}

func ptrComplex128(x []complex128) *C.doublecomplex {
	if len(x) == 0 {
		return nil
	}
	return (*C.doublecomplex)(unsafe.Pointer(&x[0]))
}

func toCInt(x []int) []C.integer {
	if len(x) == 0 {
		return nil
	}
	y := make([]C.integer, len(x))
	for i, xi := range x {
		y[i] = C.integer(xi)
	}
	return y
}

func fromCInt(x []C.integer) []int {
	if len(x) == 0 {
		return nil
	}
	y := make([]int, len(x))
	for i, xi := range x {
		y[i] = int(xi)
	}
	return y
}
