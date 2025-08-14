package main

import (
	"fmt"
	"time"
)

func listing1() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func() {
			fmt.Println("listing1", i)
		}()
	}
}

func listing2() {
	s := []int{1, 2, 3}

	for _, i := range s {
		val := i
		go func() {
			fmt.Println("listing2", val)
		}()
	}
}

func listing3() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func(val int) {
			fmt.Println("listing3", val)
		}(i)
	}
}

func main() {
	listing1()
	listing2()
	listing3()
	time.Sleep(time.Second)
}
