package races

import (
	"sync"
	"sync/atomic"
)

// this will make data race happen since 2 gorountines can access i at the same time
func listing1() {
	i := 0

	go func() {
		i++
	}()

	go func() {
		i++
	}()
}

// we could resolve this by make the operation atomic
func listing2() {
	var i int64

	go func() {
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		atomic.AddInt64(&i, 1)
	}()
}

// we could use mutex to lock it for safe access
func listing3() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
	}()

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
	}()
}

// we can use a channel to avoid data race
func listing4() {
	i := 0
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch
}

// this will not have data race but will have race condition since
// we cannot determine what is the outcome of this function
// (doesnt know what routines execute first and what is the order of execution)
func listing5() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 1
	}()

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		i = 2
	}()

	_ = i
}
