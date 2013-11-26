package clap

import (
	"errors"
	"fmt"
)

func errUnknown(info int) error {
	return fmt.Errorf("lapack info non-zero: %d", info)
}

func errNonPosDims(a Const) error {
	rows, cols := a.Dims()
	if rows == 0 || cols == 0 {
		return errors.New("matrix empty")
	}
	if rows < 0 || cols < 0 {
		return fmt.Errorf("matrix dims not positive: %dx%d", rows, cols)
	}
	return nil
}

func errNonSquare(a Const) error {
	rows, cols := a.Dims()
	if rows != cols {
		return fmt.Errorf("matrix not square: %dx%d", rows, cols)
	}
	return nil
}

func eqDims(a, b Const) bool {
	m, n := a.Dims()
	p, q := b.Dims()
	return m == p && n == q
}

func errIncompat(a Const, b []complex128) error {
	return errIncompatT(a, false, b)
}

func errIncompatT(a Const, t bool, b []complex128) error {
	rows, cols := a.Dims()
	if t {
		rows, cols = cols, rows
	}
	if rows != len(b) {
		return fmt.Errorf("incompatible: %dx%d and %d", rows, cols, len(b))
	}
	return nil
}

func errInvalidArg(arg int) error {
	return fmt.Errorf("invalid arg: %d", arg)
}

func errSingular(index int) error {
	return fmt.Errorf("exactly singular: at index %d", index)
}

func errNotPosDef(index int) error {
	return fmt.Errorf("matrix not pos def: at index %d", index)
}

func errOffDiagFailConverge(index int) error {
	return fmt.Errorf("off-diag elements failed to converge to zero: %d", index)
}

func errBadShape(m, n int) error {
	return fmt.Errorf("invalid shape: %dx%d", m, n)
}

func errNotFullRank(index int) error {
	return fmt.Errorf("not full rank: at index %d", index)
}

func errFailConverge(info int) error {
	return fmt.Errorf("did not converge: info %d", info)
}
