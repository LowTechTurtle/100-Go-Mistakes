package main

import (
	"fmt"
)

func bar() error {
	return barError{}
}

type barError struct{}

func (b barError) Error() string {
	return "bar error"
}

type BarError struct {
	err error
}

func (br BarError) Error() string {
	return "bar error:" + br.err.Error()
}

func listing1() error {
	err := bar()
	if err != nil {
		// before go 1.13 we wrap error like this, create a new type and implement Error interface
		return BarError{err}
	}
	return nil
}

func listing2() error {
	err := bar()
	if err != nil {
		// wrap the error
		return fmt.Errorf("bar failed: %w", err)
	}
	return nil
}

func listing3() error {
	err := bar()
	if err != nil {
		// this would not wrap the error but change it( cant use errors.As)
		return fmt.Errorf("bar failed: %v", err)
	}
	return nil
}
