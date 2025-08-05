package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

func main() {
	s := Store{
		m: make(map[string]*Customer),
	}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	print(s.m)
}

func (s *Store) storeCustomers(customers []Customer) {
	// in newer go version, go fixed this behavior to make it less confusing
	// in previous version, go created one customer var with the same address
	// with each iteration, they store it to the same var with the same address
	// the result is every elements get the value of the last item in the slice
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		s.m[customer.ID] = &customer
	}
}

func (s *Store) storeCustomers2(customers []Customer) {
	for _, customer := range customers {
		// in newer go version, this is no longer required
		current := customer
		s.m[current.ID] = &current
	}
}

func (s *Store) storeCustomers3(customers []Customer) {
	for i := range customers {
		customer := &customers[i]
		s.m[customer.ID] = customer
	}
}

func print(m map[string]*Customer) {
	for k, v := range m {
		fmt.Printf("key=%s, value=%#v\n", k, v)
	}
}
