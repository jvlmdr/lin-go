package lapack

import "errors"

var (
	ErrNonZeroInfo = errors.New("lapack info was non-zero")
)
