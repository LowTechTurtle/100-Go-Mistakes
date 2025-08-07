package main

import "fmt"

func main() {
	s := Struct{id: "foo"}
	// this is just like with argument, method are just syntactic sugar for
	// a function that take the first args as s
	// therefore this evaluate right away, getting the value of s
	// so it will print foo
	defer s.print()
	s.id = "bar"
}

type Struct struct {
	id string
}

func (s Struct) print() {
	fmt.Println(s.id)
}
