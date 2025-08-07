package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// s1 is sliced by byte here, not by rune
	s1 := "super massive turtle"
	s2 := s1[:5]
	fmt.Println(s2)

	// it should be converted to slice of runes to slice by rune
	s3 := string([]rune(s1)[:5])
	fmt.Println(s3)
}

type store struct{}

func (s store) handleLog1(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	// here go will create a new string pointed to the same backing array
	// that still have high capacity => mem leak
	uuid := log[:36]
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog2(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	// by converting to byte and slice it, then create a new string
	// we performed a deep copy => no mem leak
	uuid := string([]byte(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog3(log string) error {
	if len(log) < 36 {
		return errors.New("log is not correctly formatted")
	}
	// we could use strings.Clone to clone a string, avoiding mem leak
	uuid := string(strings.Clone(log[:36]))
	s.store(uuid)
	// Do something
	return nil
}

func (s store) store(uuid string) {
	// ...
}
