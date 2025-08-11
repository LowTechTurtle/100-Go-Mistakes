package main

import (
	"io"
	"sync"
	"sync/atomic"
)

func normalread(r io.Reader) (int, error) {
	count := 0
	b := make([]byte, 2048)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return count, nil
			} else {
				return count, err
			}
		}
		count += task(b)
	}
}

func concurrentread(r io.Reader) (int64, error) {
	var count int64 = 0
	b := make([]byte, 2048)

	buffer := make(chan []byte, 10)

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for buff := range buffer {
				atomic.AddInt64(&count, int64(task(buff)))
			}
		}()
	}
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return count, err
			}
		}
		buffer <- b
	}
	close(buffer)
	wg.Wait()
	return count, nil
}

func task(b []byte) int {
	return len(b)
}
