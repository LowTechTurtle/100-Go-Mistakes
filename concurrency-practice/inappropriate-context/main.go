package main

import (
	"context"
	"net/http"
	"time"
)

// context with request canceled when:
// - client connection close( in this case this is good)
// - if http/2 request, when the request is canceled( this is ok)
// - when response is written back to client, not okay, we dont have a way
// to make sure that response is written before we're done publishing in this code
func handler1(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		err := publish(r.Context(), response)
		// do sth with err
		_ = err
	}()

	writeResponse(response)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		// this will fix the problem but we cant pass the value from request context
		err := publish(context.Background(), response)
		// do sth with err
		_ = err
	}()

	writeResponse(response)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	go func() {
		// we copy the context in a custom context and return its val
		err := publish(detach{ctx: r.Context()}, response)
		// do sth with err
		_ = err
	}()

	writeResponse(response)
}

type detach struct {
	ctx context.Context
}

// now we implement the Context interface on detach
func (d detach) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detach) Done() (c <-chan struct{}) {
	return
}

func (d detach) Err() (err error) {
	return
}

func (d detach) Value(key any) any {
	return d.ctx.Value(key)
}

func doSomeTask(context.Context, *http.Request) (string, error) {
	return "", nil
}

func publish(context.Context, string) error {
	return nil
}

func writeResponse(string) {}
