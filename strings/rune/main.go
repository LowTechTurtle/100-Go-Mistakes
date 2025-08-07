package main

import "fmt"

func main() {
	// each rune could use 1 to 4 bytes
	s := "hello"
	fmt.Println(len(s))

	// this rune take 3 bytes
	s = "汉"
	fmt.Println(len(s))

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s)
}
