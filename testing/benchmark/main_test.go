package timer

import (
	"sync/atomic"
	"testing"
)

func BenchmarkFoo1(b *testing.B) {
	expensiveSetup()
	// reset timer after setup to avoid wrong benchmark because of setup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		functionUnderTest()
	}
}

func BenchmarkFoo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// we could stop timer and turn it back on when there is setup in a loop
		b.StopTimer()
		expensiveSetup()
		b.StartTimer()
		functionUnderTest()
	}
}

func functionUnderTest() {
}

func expensiveSetup() {
}

// int32 and int64 got the same performance but due to background process
// or other program running, one would seem to be faster than the other
//
// to solve that, use benchstat to run many tests and average out to results
// or just increase the run time to a bit longer
func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}
