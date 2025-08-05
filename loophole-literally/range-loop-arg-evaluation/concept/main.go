package main

func main() {
	// s1 only evaluated once so this run normally
	s1 := []int{0, 1, 2}
	for range s1 {
		s1 = append(s1, 10)
	}

	// the len(s2) evaluated everytime => infinite loop
	s2 := []int{0, 1, 2}
	for i := 0; i < len(s2); i++ {
		s2 = append(s2, 10)
	}
}
