package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	ID         string
	Operations []float32
}

func main() {
	var s1 []float32
	s2 := make([]float32, 0)

	customer1 := customer{
		ID:         "1",
		Operations: s1,
	}

	customer2 := customer{
		ID:         "2",
		Operations: s2,
	}

	// json marshal discriminate between nil and empty slice
	fmt.Println(customer1)
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	fmt.Println(customer2)
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b))
}
