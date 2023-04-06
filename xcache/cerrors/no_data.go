package cerrors

import (
	"errors"
	"fmt"
)

func NewNoDataError(key string) error {
	return &NoDataError{
		key: key,
	}
}

func IsNoDataError(in error) (err *NoDataError, ok bool) {
	if errors.As(in, &err) {
		return err, true
	}

	return nil, false
}

type NoDataError struct {
	key string
}

func (e *NoDataError) Error() string {
	return fmt.Sprintf("No data found on key '%s'", e.key)
}
