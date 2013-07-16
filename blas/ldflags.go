package blas

// #cgo linux LDFLAGS: -lblas
// #cgo darwin LDFLAGS: -framework Accelerate
import "C"
