package main

import (
	"context"
	"log"
	"time"
)

type Event struct{}

func handle(event Event) {}

func consumer1(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
			// would create a channel every loop -> leak mem
		case <-time.After(time.Hour):
			log.Println("warning: no event in an hour")
		}
	}
}

func consumer2(ch <-chan Event) {
	for {
		// create context every loop, inefficent
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no event in an hour")
		}
	}
}

func consumer3(ch <-chan Event) {
	timer := time.NewTimer(time.Hour)
	for {
		// reset the timer every loop
		timer.Reset(time.Hour)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C: // timer expiration
			log.Println("warning: no event in an hour")
		}
	}
}
