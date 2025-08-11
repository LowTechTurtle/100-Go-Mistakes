package main

import "fmt"

func listing1() {
	// create goroutines happens before goroutines begin
	// => i is safe
	i := 0
	go func() {
		i++
	}()
}

func listing2() {
	// exit of gorountines doesnt guarantee anything
	// i maybe 0 or 1
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}

func listing3() {
	// send on a channel happens before the corresponding receive on the chan
	// i++ is before send => this is safe
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	ch <- struct{}{}
}

func listing4() {
	// close of a channel happen before the receive of this closure
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	close(ch)
}

func listing5() {
	// receive from unbuffered channel happens before send complete
	// this chan is buffered, this will have data race tho
	i := 0
	ch := make(chan struct{}, 1)
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

func listing6() {
	// but this is okay and safe
	i := 0
	ch := make(chan struct{})
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}
