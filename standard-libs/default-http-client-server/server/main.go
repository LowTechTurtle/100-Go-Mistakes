package main

import (
	"net/http"
	"time"
)

func main() {
	// tls handshake -> read header -> read body -> response
	s := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond, //tls handshake , read header and body timeout
		// specify how long handler should take, counting from read body
		Handler: http.TimeoutHandler(handler{}, time.Second, "foo"),
	}
	_ = s
}

type handler struct{}

func (h handler) ServeHTTP(http.ResponseWriter, *http.Request) {}
