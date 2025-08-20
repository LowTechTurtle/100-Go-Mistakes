package main

import (
	"sync"
	"time"
)

type Event struct {
	Timestamp time.Time
	Data      string
}

type Cache struct {
	mu     sync.RWMutex
	events []Event
}

func (c *Cache) TrimOlderThan(since time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// using time.Now can lead to flaky test

	t := time.Now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
		}
	}
}

func (c *Cache) Add(events []Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events = append(c.events, events...)
}

func (c *Cache) GetAll() []Event {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.events
}
