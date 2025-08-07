package main

import "fmt"

func main() {
	s := &Struct{id: "foo"}
	// this is just like with argument, method are just syntactic sugar for
	// a function that take the first args as s
	// therefore this evaluate right away, getting the address of s
	// after this the change will be reflected because we copied address
	defer s.print()
	s.id = "bar"
}

type Struct struct {
	id string
}

func (s *Struct) print() {
	fmt.Println(s.id)
}
