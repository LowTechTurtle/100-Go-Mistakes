package main

import (
	"fmt"

	_ "github.com/LowTechTurtle/100-Go-Mistakes/code-proj-organize/misuse-init/redis"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("in main")
}
