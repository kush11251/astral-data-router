package errors

import (
	"errors"
)

func NewError(message string) error {
	return errors.New(message)
}
