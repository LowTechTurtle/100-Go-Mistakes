package main

import "fmt"

type data struct {
	id int
	balance *float64
}

func (d data) add(operation float64) {
	d.id += 1
	*d.balance += operation
}

func main() {
	var x float64 = 100
	c := data{1, &x}
	// since we copied the struct => we copied the address of the pointer in the struct
	// we change the value at that address => the value of the var pointed by the pointer will be changed
	c.add(50.)
	fmt.Printf("id: %d, balance: %.2f\n", c.id, *c.balance)
}
