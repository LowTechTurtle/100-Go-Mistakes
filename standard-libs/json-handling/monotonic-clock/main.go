package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Time time.Time
}

func listing1() error {
	// time usually have monotonic time at the end
	t := time.Now()
	event1 := Event{t}

	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event
	// when unmarshal, time normally lost the monotonic part
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
	fmt.Println(event1.Time == event2.Time)
	fmt.Println(event1.Time.Equal(event2.Time))
	return nil
}

func listing2() error {
	t := time.Now()
	event1 := Event{t.Truncate(0)} // truncate the monotonic time

	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
	fmt.Println(event1.Time == event2.Time)
	return nil
}

func main() {
	listing1()
	listing2()
}
