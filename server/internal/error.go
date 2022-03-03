package internal

import (
	"errors"
	"fmt"
)

func NewError(message string, err error) error {
	text := fmt.Sprintf("%v: %v", message, err)
	return errors.New(text)
}
