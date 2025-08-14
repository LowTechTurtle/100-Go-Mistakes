package main

import "fmt"

func listing1() {
	s := make([]int, 1)

	go func() {
		s1 := append(s, 1)
		// this doesnt have data race because s1 have len = 1, cap =1
		// when called append, it will make a copy => no data race
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()
}

func listing2() {
	s := make([]int, 0, 1)

	go func() {
		// this will have data race because len = 0, cap = 1
		// it will try to write the existing free space => data race
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s1 := append(sCopy, 1)
		fmt.Println(s1)
	}()

	go func() {
		sCopy := make([]int, len(s), cap(s))
		copy(sCopy, s)

		s2 := append(sCopy, 1)
		fmt.Println(s2)
	}()
}
