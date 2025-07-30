package main

import "fmt"

// doest solve port management( eg. zero value is 0 and unset is 0)
// => need pointer to do it but if pointer is used in config struct
// it make the api hard to use and navigate around
type Config struct {
	Port int
}

// stub func
func NewServer(config Config) {
	fmt.Println(config.Port)
}

func main() {
	NewServer(Config{80})
}
