package main

type node struct {
	value int64
	next  *node
}

// store value and pointer, this is a non-unit stride, cpu can't predict where
// is the next piece of data so it runs slower
func linkedList(n *node) int64 {
	var total int64
	for n != nil {
		total += n.value
		n = n.next
	}
	return total
}

// this is every 2 elements, non-unit stride, runs slower than unit stride(every element)
// it is still predictable so it is remakably faster than linked list
func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}
