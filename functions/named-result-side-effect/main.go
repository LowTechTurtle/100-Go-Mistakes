package main

import (
	"context"
	"errors"
)

type loc struct{}

func (l loc) validateAddress(address string) bool {
	return true
}

func (l loc) getCoordinates1(ctx context.Context, address string) (
	lat, long float64, err error) {
	if !l.validateAddress(address) {
		return 0, 0, errors.New("invalid address")
	}

	if ctx.Err() != nil {
		// this will return err zero value stated above in the named return val
		// => bug
		return 0, 0, err
	}

	// get and return addr
	return
}

func (l loc) getCoordinates2(ctx context.Context, address string) (
	lat, long float64, err error) {
	if !l.validateAddress(address) {
		return 0, 0, errors.New("invalid address")
	}

	if err := ctx.Err(); err != nil {
		// fixed by doing an assignment to err, this will shadow err but will work
		// as intended
		return 0, 0, err
	}

	// get and return addr
	return
}

func (l loc) getCoordinates3(ctx context.Context, address string) (
	lat, long float64, err error) {
	if !l.validateAddress(address) {
		return 0, 0, errors.New("invalid address")
	}

	if err = ctx.Err(); err != nil {
		// this blank statement will return error and not shadow anything
		// but it is not idiomatic to use naked return mix with arguments return
		// it makes the code confusing and hard to read, and might make
		// delicate bugs
		return
	}

	// get and return addr
	return
}
