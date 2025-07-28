package shadowing

import (
	"log"
	"net/http"
)

func ShadowingFix2() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	//use client in some logic
	log.Println(client)
	return nil
}
