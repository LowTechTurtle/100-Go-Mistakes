package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Cache{
		balances: make(map[string]float64),
	}
	c.AddBalance("1", 1.0)
	c.AddBalance("2", 3.0)
	fmt.Println(c.AverageBalance1())
	fmt.Println(c.AverageBalance2())
	fmt.Println(c.AverageBalance3())
}

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance1() float64 {
	c.mu.RLock()
	// this will only create a reference to the map, not copying the map itself
	balances := c.balances
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances {
		// => data race
		sum += balance
	}
	return sum / float64(len(balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	sum := 0.
	// this will not has data race because we locked it with mutex
	// if the operation on map is short, we could handle it like this
	// if the operation is longer, we should not hold the lock
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance3() float64 {
	c.mu.RLock()
	// we could copy the map like so if the opration after copying the map is
	// complicated and need time, otherwise, this is prove marginal time using the lock
	// but twice operation needed O(2n)
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}
