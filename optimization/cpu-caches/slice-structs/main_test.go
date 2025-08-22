package main

import "testing"

var global int64

const n = 1_000_000

// this cache line would have 1 a element followed by 1 b element, and then 1a, 1b,...
func BenchmarkSumFoo(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]Foo, n)
		b.StartTimer()
		local = sumFoo(s)
	}
	global = local
}

// this will have a line of a element then a line of b element
// this would need only half of the cache lines so it is faster
func BenchmarkSumBar(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		bar := Bar{
			a: make([]int64, n),
			b: make([]int64, n),
		}
		b.StartTimer()
		local = sumBar(bar)
	}
	global = local
}
