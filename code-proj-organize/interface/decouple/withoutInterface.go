package main

import "fmt"

// stub external data store
type Store struct {
	key string
	val string
}

type DataStore struct {
	store Store // depend directly on it
}

func (ds *DataStore) Store(key, val string) error {
	ds.store.key = key
	ds.store.val = val
	return nil
}

func (ds *DataStore) Get(key string) (string, error) {
	if key == ds.store.key {
		return ds.store.val, nil
	}
	return "", fmt.Errorf("key not found")
}
