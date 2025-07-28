package shadowing

import (
	"log"
	"net/http"
)

func ShadowingFix1() error {
	var client *http.Client
	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		client = c

	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	//use client in some logic
	log.Println(client)
	return nil
}
