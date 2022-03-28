package internal

import (
	"errors"
	"fmt"
	"io"
)

func NewError(message string, err error) error {
	text := fmt.Sprintf("%v: %v", message, err)
	return errors.New(text)
}

func CloseOrPanic(closer io.Closer) func() {
	return func() {
		if err := closer.Close(); err != nil {
			text := fmt.Sprintf("Close error: %v", err)
			panic(text)
		}
	}
}
