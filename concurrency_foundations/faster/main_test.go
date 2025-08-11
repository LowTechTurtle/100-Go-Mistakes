package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMergeSortV1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandom()
		b.StartTimer()
		MergeSortV1(input)
	}
}

func BenchmarkMergeSortV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandom()
		b.StartTimer()
		MergeSortV2(input)
	}
}

func getRandom() []int {
	n := 10_000
	res := make([]int, n)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	for i := 0; i < n; i++ {
		res[i] = rnd.Int()
	}
	return res
}
