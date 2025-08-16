package main

import "encoding/json"

func listing1() error {
	b := getMessage()
	var m map[string]any // every numerical is converted to float64
	// even if it has no floating point in it
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	return nil
}

func getMessage() []byte {
	return nil
}
