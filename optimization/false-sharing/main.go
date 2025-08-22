package main

import "sync"

type Input struct {
	a int64
	b int64
}

type Result1 struct {
	sumA int64
	sumB int64
}

// normally sumA and sumB is allocated next to each other on the cache line
// so they are on different L1 cache on different cores
// when it update in one, the other must invalidate the cache
// hence very low speed
func count1(inputs []Input) Result1 {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result1{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}

type Result2 struct {
	sumA int64
	_    [56]byte
	sumB int64
}

// but when we add a padding so that sumB is on another cache line on L1
// when we update sumA or sumB, it doesnt need to invalidate the cache in the other core
// so it could keep using the mem in the cache => execute twice as fast
func count2(inputs []Input) Result2 {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result2{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()
	return result
}
