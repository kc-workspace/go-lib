package dterrors

import (
	"errors"
	"fmt"
)

func NewConvertFailError(input any, outType string) error {
	return &ConvertFailError{
		input:   input,
		outType: outType,
	}
}

func IsConvertFailError(in error) (err *ConvertFailError, ok bool) {
	if errors.As(in, &err) {
		return err, true
	}

	return nil, false
}

type ConvertFailError struct {
	input   any
	outType string
}

func (e *ConvertFailError) Error() string {
	return fmt.Sprintf("Cannot convert '%v' to %s", e.input, e.outType)
}
