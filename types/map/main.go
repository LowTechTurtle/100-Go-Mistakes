package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Init
	n := 1_000_000
	// can use pointer instead of [128]bute
	// it still leak mem but instead of leak 128 for each entry, now it only leak 4

	// if we need to store more than 128 byte, golang will automagically store pointer
	// make it leak as much as using a pointer here
	m := make(map[int]*[128]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// End
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func randBytes() *[128]byte {
	return &[128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
