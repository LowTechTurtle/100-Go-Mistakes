package main

import "fmt"

func main() {
	var a float64
	// 3 special float type: negative infinite, positive infinite and NaN
	neginf := -1 / a
	posinf := 1 / a
	nan := a / a
	fmt.Println(neginf, posinf, nan)

	// reminder: floating number does not save the precise number
	// so when comparing, we should check if their difference is in an acceptable range

	f1()
	f2()
	f3()
}

// when adding float, add the group with the same magnitude for better accuracy (f2 is better than f1)
func f1() {
	var a float64 = 100000
	for i := 0; i < 100000; i++ {
		a += 1.001
	}
	fmt.Println(a)
}

func f2() {
	var a float64
	for i := 0; i < 100000; i++ {
		a += 1.001
	}
	a += 100000
	fmt.Println(a)
}

// do multiplying before adding or subtracting for better accuracy
func f3() {
	a := 100000.001
	b := 1.0001
	c := 1.0002
	fmt.Println(a * (b + c))
	fmt.Println(a*b + a*c)
}
