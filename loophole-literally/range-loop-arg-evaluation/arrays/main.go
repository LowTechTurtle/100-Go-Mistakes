package main

import "fmt"

func listing1() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10
		// the value already copied to memory, a[2] is 10 but v is still 2( illustrated in listing2)
		if i == 2 {
			fmt.Println(v)
		}
	}
}

func listing2() {
	a := [3]int{0, 1, 2}
	for i := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(a[2])
		}
	}
}

func listing3() {
	a := [3]int{0, 1, 2}
	// here we will deference and read from address
	// go copied the address
	// when read v, go read the value in the address( stored in v) so it can read 10
	for i, v := range &a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
}

func main() {
	listing1()
	listing2()
	listing3()
}
