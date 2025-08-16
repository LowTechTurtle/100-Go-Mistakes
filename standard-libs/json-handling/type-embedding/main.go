package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	if err := listing1(); err != nil {
		panic(err)
	}
	if err := listing2(); err != nil {
		panic(err)
	}
	if err := listing3(); err != nil {
		panic(err)
	}
}

type Event1 struct {
	ID int
	time.Time
}

func listing1() error {
	event := Event1{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// this doesnt have ID because time.Time is embeded, and it
	// implement a custom marshal method, since its embeded, that method is promoted
	// hence, the weird behavior
	fmt.Println(string(b))
	return nil
}

type Event2 struct {
	ID   int
	Time time.Time
}

func listing2() error {
	event := Event2{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// by not using embedding, the marshal method works as intended
	fmt.Println(string(b))
	return nil
}

type Event3 struct {
	ID int
	time.Time
}

// create a custom marshal method
func (e Event3) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			ID   int
			Time time.Time
		}{
			ID:   e.ID,
			Time: e.Time,
		},
	)
}

func listing3() error {
	event := Event3{
		ID:   1234,
		Time: time.Now(),
	}

	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// we can still embed and get the marshal work correctly by defining our
	// custom marshal method
	fmt.Println(string(b))
	return nil
}
