package main

import (
	"net"
	"net/http"
	"time"
)

func main() {
	// we shouldnt use the default client as it offer no timeout
	client := &http.Client{
		Timeout: 5 * time.Second, // global timeout( total time for all operation: dial,
		// tls handshake, request, resp header read, resp body read)
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second, // establish TCP connection timeout
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second, // read the resp header timeout
		},
	}
	_ = client
}
