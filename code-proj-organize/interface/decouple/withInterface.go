package main

import "fmt"

// create an abstraction
type Storer interface {
	Store(string, string) error
	Get(string) (string, error)
}

// does not depend on the actual implementation, instead, it describe what i need
type BetterDataStore struct {
	store Storer
}

func (bds BetterDataStore) Store(key, val string) error {
	bds.store.Store(key, val)
	return nil
}

func (bds BetterDataStore) Get(key string) (string, error) {
	return bds.store.Get(key)
}

// we can just change the factory function, then we can swap between implementations easily
func NewDataStore() BetterDataStore {
	return BetterDataStore{store: &DataStore{}}
}

func main() {
	myData := NewDataStore()
	myData.Store("banana", "turtle")
	val, _ := myData.Get("banana")
	fmt.Println(val)
}
