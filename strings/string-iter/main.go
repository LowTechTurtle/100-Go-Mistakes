package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))

	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}

	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}

	s2 := "hello"
	fmt.Printf("%c\n", rune(s2[4]))

	r, _ := RuneAt("hêllô", 1)
	fmt.Printf("%c", r)
}

func RuneAt(s string, i int) (rune, error) {
	var zerorune rune
	if i >= utf8.RuneCountInString(s) {
		return zerorune, errors.New("index larger or equal than rune length")
	}
	for j, r := range s {
		if j == i {
			return r, nil
		}
	}
	return zerorune, nil
}
