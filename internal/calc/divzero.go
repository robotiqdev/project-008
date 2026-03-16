package calc

import calcerrors "github.com/repo/calculator/internal/errors"

func validateDivisor(b float64) error {
	if b == 0.0 {
		return calcerrors.ErrDivByZero
	}
	return nil
}
