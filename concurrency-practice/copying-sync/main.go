package main

import (
	"sync"
	"time"
)

func main() {
	counter := NewCounter()

	go func() {
		counter.Increment1("foo")
	}()
	go func() {
		counter.Increment1("bar")
	}()

	time.Sleep(10 * time.Millisecond)
}

type Counter struct {
	mu       sync.Mutex // bad
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

func (c Counter) Increment1(name string) {
	// this will pass the lock by value, not actually locking the crit section( shared section)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func (c *Counter) Increment2(name string) {
	// passing the copy of the address will no longer have the problem
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

type Counter2 struct {
	mu *sync.Mutex // or using pointer for Mutex would be a solution for not using pointer reciver
	// however, it will require creating and &sync.Mutex{} for factory method because
	// the zero value for pointer is nil
	counters map[string]int
}

func NewCounter2() Counter2 {
	return Counter2{
		mu:       &sync.Mutex{},
		counters: map[string]int{},
	}
}
