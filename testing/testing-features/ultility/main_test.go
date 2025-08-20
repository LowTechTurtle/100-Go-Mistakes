package main

import (
	"errors"
	"testing"
)

func TestCustomer1(t *testing.T) {
	customer, err := createCustomer1("foo")
	if err != nil {
		t.Fatal(err)
	}
	// ...
	_ = customer
}

func createCustomer1(someArg string) (Customer, error) {
	customer, err := customerFactory(someArg)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func TestCustomer2(t *testing.T) {
	customer := createCustomer2(t, "foo")
	// ...
	_ = customer
}

// we can simplify err handling by passing *Testing.T and call it in the ultility function
func createCustomer2(t *testing.T, someArg string) Customer {
	customer, err := customerFactory(someArg)
	if err != nil {
		t.Fatal(err)
	}
	return customer
}

func customerFactory(someArg string) (Customer, error) {
	if someArg == "" {
		return Customer{}, errors.New("empty")
	}
	return Customer{id: someArg}, nil
}
