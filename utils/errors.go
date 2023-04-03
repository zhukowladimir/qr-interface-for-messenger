package utils

import (
	"errors"
	"fmt"
	"log"
)

type basicDeferrable func() error

func WrapError(f basicDeferrable) {
	err := f()
	if err != nil {
		log.Println("encountered error:", err.Error())
	}
}

func WrapFail(whatFailed string, err error) error {
	if err == nil {
		return nil
	}

	return errors.New(fmt.Sprintf("couldn't %s: %s", whatFailed, err.Error()))
}
