package cmat

import (
	"errors"
	"fmt"
)

func errIfDimsNotEq(a, b dimser) error {
	if !eqDims(a, b) {
		return errDimsNotEq(a, b)
	}
	return nil
}

func eqDims(a, b dimser) bool {
	m, n := a.Dims()
	p, q := b.Dims()
	return m == p && n == q
}

func errDimsNotEq(a, b dimser) error {
	m, n := a.Dims()
	p, q := b.Dims()
	return fmt.Errorf("different dims: %dx%d, %dx%d", m, n, p, q)
}

func errMulDimsMat(a, b dimser) error {
	m, n := a.Dims()
	p, q := b.Dims()
	return errMulDims(m, n, p, q)
}

func errMulDimsVec(a dimser, b []complex128) error {
	m, n := a.Dims()
	p := len(b)
	return errMulDims(m, n, p, 1)
}

func errMulDims(m, n, p, q int) error {
	if n != p {
		return fmt.Errorf("incompatible dims: %dx%d, %dx%d", m, n, p, q)
	}
	return nil
}

func errRagged(first, other int) error {
	return fmt.Errorf("ragged list of arrays: found %d and %d", first, other)
}

func errRectOutsideMat(r Rect, m, n int) error {
	rect := fmt.Sprintf("(%d, %d)-(%d, %d)", r.Min.I, r.Min.J, r.Max.I, r.Max.J)
	size := fmt.Sprintf("%dx%d", m, n)
	detail := rect + " not in " + size
	return errors.New("rectangle outside bounds: " + detail)
}
