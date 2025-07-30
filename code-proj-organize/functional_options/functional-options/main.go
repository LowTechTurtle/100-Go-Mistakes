package main

import (
	"errors"
	"fmt"
	"net/http"
)

const defaultHTTPPort = 8080

type options struct {
	port *int
}

func randomPort() int {
	return 1234
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options

	// building the options struct
	for _, f := range opts {
		err := f(&options)
		if err != nil {
			return nil, err
		}
	}

	// we manange the port accordingly to the options struct
	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}
	fmt.Println(port)
	return nil, nil
}

func client() {
	_, _ = NewServer("localhost", WithPort(8080))
}

func main() {
	client()
}
