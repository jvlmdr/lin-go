package lapack

import (
	"github.com/jackvalmadre/lin-go/mat"
	"math"
)

// Projects a vector on to the (right) null-space of a matrix.
// Computes x - A' (A A')^{-1} A x.
//
// Returns an error if matrix is not full rank.
func projNull(a Const, x []float64) ([]float64, error) {
	y, err := projRow(a, x)
	if err != nil {
		return nil, err
	}
	return minus(x, y), nil
}

// Projects a vector on to the row-space of a matrix.
// Computes A' (A A')^{-1} A x.
//
// Returns an error if matrix is not full rank.
func projRow(a Const, x []float64) ([]float64, error) {
	// y <- (A A')^{-1} A x
	y, err := SolveFullRank(mat.T(a), x)
	if err != nil {
		return nil, err
	}
	// y <- A' y
	y = mat.MulVec(mat.T(a), y)
	return y, nil
}

func plus(a, b []float64) []float64 {
	c := make([]float64, len(a))
	for i := range c {
		c[i] = a[i] + b[i]
	}
	return c
}

func minus(a, b []float64) []float64 {
	c := make([]float64, len(a))
	for i := range c {
		c[i] = a[i] - b[i]
	}
	return c
}

func scale(k float64, x []float64) []float64 {
	y := make([]float64, len(x))
	for i := range x {
		y[i] = k * x[i]
	}
	return y
}

func norm(a []float64) float64 {
	return math.Sqrt(sqrnorm(a))
}

func sqrnorm(a []float64) float64 {
	var t float64
	for _, x := range a {
		t += x * x
	}
	return t
}
