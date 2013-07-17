package zvec

import (
	"github.com/jackvalmadre/lin-go/vec"
	"math/cmplx"
)

// Compute complex conjugate for every element in the vector.
// Lazily evaluated.
func Conj(x Const) Const {
	return Map(x, cmplx.Conj)
}

// Extract real part of every element in the vector.
// Lazily evaluated.
func Real(x Const) vec.Const {
	f := func(x complex128) float64 { return real(x) }
	return MapToReal(x, f)
}

// Extract imaginary part of every element in the vector.
// Lazily evaluated.
func Imag(x Const) vec.Const {
	f := func(x complex128) float64 { return imag(x) }
	return MapToReal(x, f)
}

// Compose complex vector from real and imaginary parts.
// Lazily evaluated.
func Complex(x, y vec.Const) Const {
	f := func(x, y float64) complex128 { return complex(x, y) }
	return MapToComplex(x, y, f)
}

// Compute |x| for every element in the vector.
// Lazily evaluated.
func Abs(x Const) vec.Const {
	return MapToReal(x, cmplx.Abs)
}

// Compute exp(x) for every element in the vector.
// Lazily evaluated.
func Exp(x Const) Const {
	return Map(x, cmplx.Exp)
}

// Compute log(x) for every element in the vector.
// Lazily evaluated.
func Log(x Const) Const {
	return Map(x, cmplx.Log)
}

// Compute sqrt(x) for every element in the vector.
// Lazily evaluated.
func Sqrt(x Const) Const {
	return Map(x, cmplx.Sqrt)
}

// Compute x^p for every element in the vector.
// Lazily evaluated.
func Pow(x Const, p complex128) Const {
	f := func(x complex128) complex128 { return cmplx.Pow(x, p) }
	return Map(x, f)
}

// Vector whose entries are random and normally distributed.
func Randn(n int) Const {
	x := vec.Randn(n)
	return Complex(x, x)
}

// Applies a unary function which maps complex to real.
// Lazily evaluated.
func MapToReal(x Const, f func(complex128) float64) vec.Const {
	return mapToRealExpr{x, f}
}

type mapToRealExpr struct {
	X Const
	F func(complex128) float64
}

func (expr mapToRealExpr) Len() int         { return expr.X.Len() }
func (expr mapToRealExpr) At(i int) float64 { return expr.F(expr.X.At(i)) }

// Mutable wrapper which accesses the real part of a complex vector.
func RealMutable(x Mutable) vec.Mutable { return realExpr{x} }

type realExpr struct{ X Mutable }

func (expr realExpr) Len() int         { return expr.X.Len() }
func (expr realExpr) At(i int) float64 { return real(expr.X.At(i)) }

func (expr realExpr) Set(i int, x float64) {
	expr.X.Set(i, complex(x, imag(expr.X.At(i))))
}

// Mutable wrapper which accesses the imaginary part of a complex vector.
func ImagMutable(x Mutable) vec.Mutable { return imagExpr{x} }

type imagExpr struct{ X Mutable }

func (expr imagExpr) Len() int         { return expr.X.Len() }
func (expr imagExpr) At(i int) float64 { return imag(expr.X.At(i)) }

func (expr imagExpr) Set(i int, x float64) {
	expr.X.Set(i, complex(real(expr.X.At(i)), x))
}

// Mutable wrapper which accesses every element as its conjugate.
func ConjMutable(x Mutable) Mutable {
	return conjExpr{x}
}

type conjExpr struct{ X Mutable }

func (expr conjExpr) Len() int            { return expr.X.Len() }
func (expr conjExpr) At(i int) complex128 { return cmplx.Conj(expr.X.At(i)) }

func (expr conjExpr) Set(i int, x complex128) {
	expr.X.Set(i, cmplx.Conj(x))
}

// Applies a binary function which maps real to complex.
// Lazily evaluated.
func MapToComplex(x, y vec.Const, f func(float64, float64) complex128) Const {
	return mapToComplex{x, y, f}
}

type mapToComplex struct {
	X vec.Const
	Y vec.Const
	F func(float64, float64) complex128
}

func (expr mapToComplex) Len() int { return expr.X.Len() }

func (expr mapToComplex) At(i int) complex128 {
	return expr.F(expr.X.At(i), expr.Y.At(i))
}
