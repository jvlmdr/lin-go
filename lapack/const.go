package lapack

import "fmt"

func ErrNonZeroInfo(info int) error {
	return fmt.Errorf("lapack info non-zero: %d", info)
}
