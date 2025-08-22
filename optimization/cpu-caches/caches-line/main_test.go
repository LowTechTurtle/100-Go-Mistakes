package main

import "testing"

var global int64

func BenchmarkSum2(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum2(s)
	}
	global = local
}

func BenchmarkSum16(b *testing.B) {
	var local int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := make([]int64, 1_000_000)
		b.StartTimer()
		local = sum16(s)
	}
	global = local
}
