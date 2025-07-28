package shadowing

import (
	"log"
	"net/http"
)

// a stub variable to know if tracing is enabled
var tracing bool

// stub func
func createClientWithTracing() (*http.Client, error) {
	var c *http.Client
	var err error
	return c, err
}

// stub func
func createDefaultClient() (*http.Client, error) {
	var c *http.Client
	var err error
	return c, err
}

func wrongShadowing() error {
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing() // shadowed the client var in this block with :=
		// as a result, the client var in the outer block is always nil( zero value for pointer)
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	//use client in some logic
	log.Println(client)
	return nil
}
