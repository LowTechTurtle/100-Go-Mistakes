package main

import (
	"io"
	"log"
	"net/http"
)

func (h handler) getBody1() (string, error) {
	// forgot to close the response body
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (h handler) getBody2() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	return string(body), nil
}

func (h handler) getStatusCode1(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	// if we just want the status code, not resp body
	// we still need to close it anyway
	// but since we're not reading it and we close it
	// the connection will usually be terminated
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	return resp.StatusCode, nil
}

func (h handler) getStatusCode2(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	// we should close the response body when there are no error
	// also, every types that implement closer method should be closed
	// when we're done with it to prevent leak
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	// we use io.Copy to io.Discard to read the body and get rid of it
	// its faster than ReadAll and keeps the connection alive
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}

type handler struct {
	client http.Client
	url    string
}
