package main

import (
	"fmt"
	"sync"
	"time"
)

func listing1() {
	type Donation struct {
		mu      sync.RWMutex
		balance int
	}

	donation := &Donation{}

	f := func(goal int) {
		donation.mu.RLock()
		// busy waiting, bad approach
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
}

func listing2() {
	type Donation struct {
		balance int
		ch      chan int
	}

	donation := &Donation{ch: make(chan int)}

	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}
	go f(10)
	go f(15)
	// thiss wont work because message sent on a channel with be read by only
	// one go routines => f(10) may print 11 other than 10
	for {
		time.Sleep(100 * time.Millisecond)
		donation.balance++
		donation.ch <- donation.balance
	}
}

func listing3() {
	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{cond: sync.NewCond(&sync.Mutex{})}

	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait() // this will unlock the mutex lock and suspend this goroutine
			// until a broadcast or a signal, then it will lock the mutex again
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}

	go f(10)
	go f(15)

	for {
		time.Sleep(100 * time.Millisecond)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast() // broadcast to unlock the mutex, check the condition
	}
}

func main() {
	listing3()
}
