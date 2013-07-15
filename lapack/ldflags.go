package lapack

// #cgo linux LDFLAGS: -llapack -lblas
// #cgo darwin LDFLAGS: -framework Accelerate
import "C"
