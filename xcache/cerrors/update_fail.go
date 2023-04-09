package cerrors

import (
	"errors"
	"fmt"
)

func NewUpdateFailError(key string) error {
	return &UpdateFailError{
		key: key,
	}
}

func IsUpdateFailError(in error) (err *UpdateFailError, ok bool) {
	if errors.As(in, &err) {
		return err, true
	}

	return nil, false
}

type UpdateFailError struct {
	key string
}

func (e *UpdateFailError) Error() string {
	return fmt.Sprintf("cannot update key '%s'", e.key)
}
