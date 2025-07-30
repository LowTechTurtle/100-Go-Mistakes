package main

import (
	"errors"
	"fmt"
	"net/http"
)

type Config struct {
	Port int
}

// totally fair and randomized port
func randomPort() int {
	return 1234
}

const defaultHTTPPort = 8080

// now we use the port *int to know if 0 is unset or set at 0
type ConfigBuilder struct {
	port *int
}

// method expose for user to choose port
// return the builder so that we can chain call
func (cb *ConfigBuilder) Port(p int) *ConfigBuilder {
	cb.port = &p
	return cb
}

// main logic for port management and create config struct
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = defaultHTTPPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}

	return cfg, nil
}

func NewServer(addr string, config Config) (*http.Server, error) {
	fmt.Println(config.Port)
	return nil, nil
}

// in short, instead of directly use the struct, we can create a builder struct
// define method for each config, then invoke build to run management logic
// and return the config struct
func main() {
	builder := ConfigBuilder{}
	builder.Port(80)
	cfg, err := builder.Build()
	if err != nil {
		panic(err)
	}
	NewServer("localhost", cfg)
}
