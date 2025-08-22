package main

const n = 1_000_000

// this need to read s[0], increment it, then read s[0], get the remainder
// then read and increment s[1]
func add(s [2]int64) [2]int64 {
	for i := 0; i < n; i++ {
		s[0]++
		if s[0]%2 == 0 {
			s[1]++
		}
	}
	return s
}

// this read s[0], increment it and get the remainder at the same time, and
// then read and increment s[1], it add cpu-level parallelism so it can execute faster
//
// the first function has more data and control hazard( the latter instruction
// must wait for the previous instruction execute and evaluate the previous var)
//
// to avoid this, we should try to write code so that the later depend on the previous line
// as little as possible
func add2(s [2]int64) [2]int64 {
	for i := 0; i < n; i++ {
		v := s[0]
		s[0] = v + 1
		if v%2 != 0 {
			s[1]++
		}
	}
	return s
}
