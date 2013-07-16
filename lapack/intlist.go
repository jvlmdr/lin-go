package lapack

// #include "../f2c.h"
import "C"

// This type only exists to get as much code as possible out of CGo.
type IntList []C.integer
