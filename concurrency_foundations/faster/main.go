package main

import "sync"

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func MergeSortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	mid := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		MergeSortV1(s[:mid])
	}()

	go func() {
		defer wg.Done()
		MergeSortV1(s[mid:])
	}()
	wg.Wait()
	merge(s, mid)
}

func MergeSortV2(s []int) {
	thresHold := 2048
	if len(s) <= 1 {
		return
	}

	mid := len(s) / 2

	if len(s) > thresHold {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			MergeSortV2(s[:mid])
		}()

		go func() {
			defer wg.Done()
			MergeSortV2(s[mid:])
		}()

		wg.Wait()
		merge(s, mid)
	} else {
		MergeSortV2(s[:mid])
		MergeSortV2(s[mid:])
		merge(s, mid)
	}

}
