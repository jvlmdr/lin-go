package clap

// #cgo linux LDFLAGS: -llapack -lblas
// #cgo darwin LDFLAGS: -framework Accelerate
import "C"
