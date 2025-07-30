package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// now Foo can access Baz directly because Baz was promoted
type Foo struct {
	Bar
}

type Bar struct {
	Baz int
}

func promote() {
	f := Foo{}
	f.Baz = 5
	fmt.Println(f.Bar.Baz)
}

type InMem struct {
	sync.Mutex
	m map[string]int
}

// because of promoting, we can access the method even if we don't want to export it
// what if they can call InMem.New() ? not good
func (i *InMem) Get(key string) (int, bool) {
	i.Lock()
	v, contains := i.m[key]
	i.Unlock()
	return v, contains
}

// now we can expose only the methods we want to
type InMemBetter struct {
	mu sync.Mutex
	m  map[string]int
}

// we should just embed io.Writecloser to make the call less cumbersome and
// the code easier to understand
type Logger struct {
	writeCloser io.WriteCloser
}

func (l Logger) Write(p []byte) (int, error) {
	return l.writeCloser.Write(p)
}

func (l Logger) Close() error {
	return l.writeCloser.Close()
}

func main() {
	promote()
	l := Logger{writeCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
