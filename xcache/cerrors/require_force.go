package cerrors

import (
	"errors"
	"fmt"
)

func NewRequireForceError(action, key string) error {
	return &RequireForceError{
		action: action,
		key:    key,
	}
}

func IsRequireForceError(in error) (err *RequireForceError, ok bool) {
	if errors.As(in, &err) {
		return err, true
	}

	return nil, false
}

type RequireForceError struct {
	action string
	key    string
}

func (e *RequireForceError) Error() string {
	return fmt.Sprintf("require force %s on key '%s'", e.action, e.key)
}
