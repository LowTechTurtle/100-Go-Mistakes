package main

import "fmt"

func badCopy() {
	src := []int{1, 2, 3}
	var dest []int

	copy(dest, src)
	// since dest is nil, len = 0, copy take the least length of 2 slices to copy
	// so 0 is the number of element of the slice that will be copied
	// aka no copying shite
	fmt.Println(dest)
}

func goodCopy() {
	src := []int{1, 2, 3}
	dest := make([]int, len(src))

	copy(dest, src)
	fmt.Println(dest)
}

func main() {
	badCopy()
	goodCopy()
}
