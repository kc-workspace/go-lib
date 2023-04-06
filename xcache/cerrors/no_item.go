package cerrors

import (
	"errors"
	"fmt"
)

func NewNoItemError(key string) error {
	return &NoItemError{
		key: key,
	}
}

func IsNoItemError(in error) (err *NoItemError, ok bool) {
	if errors.As(in, &err) {
		return err, true
	}

	return nil, false
}

type NoItemError struct {
	key string
}

func (e *NoItemError) Error() string {
	return fmt.Sprintf("No item found on key '%s'", e.key)
}
