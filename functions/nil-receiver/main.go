package main

import (
	"errors"
	"log"
	"strings"
)

type MultiError struct {
	errs []string
}

func (m *MultiError) Add(err error) {
	m.errs = append(m.errs, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errs, ";")
}

type Customer struct {
	Age  int
	Name string
}

// mistake: returning a nil pointer receiver
// when returning a nil MultiError pointer, it got wrapped by
// the error interface, the wrapee is nil but wrapper is not nil
// => always return a non nil error
func (c Customer) Validate1() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	return m
}

func (c Customer) Validate2() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	if m != nil {
		return m
	}
	return nil
}

func main() {
	customer := Customer{Age: 33, Name: "John"}
	if err := customer.Validate1(); err != nil {
		log.Fatalf("customer is invalid: %v", err)
	}
}
