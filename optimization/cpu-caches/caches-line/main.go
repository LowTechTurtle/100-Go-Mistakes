package main

// this function still use half of its cache line so its pretty fast
func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}

// theoretically, this should be 8 times faster than sum2
// but it only about 50% faster, because everytime
// it needs a number, it must get it from another cache line
func sum16(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 16 {
		total += s[i]
	}
	return total
}
