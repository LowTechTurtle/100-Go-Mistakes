package main

func listing1() {
	a := 3
	b := 2

	c := sumValue(a, b)
	println(c)
}

//go:noinline
func sumValue(x, y int) int {
	z := x + y
	return z
}

func listing2() {
	a := 3
	b := 2

	c := sumPtr(a, b)
	println(*c)
}

//go:noinline
func sumPtr(x, y int) *int {
	z := x + y
	// this z variable must escape to heap because if it lives on the stack
	// it will be unacessible after this function exit
	// if we could even access at that specific memory, we could've long overwrite
	// it with other values in our stack
	return &z
}

func listing3() {
	a := 3
	b := 2
	c := sum(&a, &b)
	println(c)
}

//go:noinline
func sum(x, y *int) int {
	return *x + *y
}
