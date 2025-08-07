package main

import (
	"bytes"
	"io"
	"strings"
)

// we should favor using operation in bytes package instead of convert bytes to string
// and convert it back
// most data send and received in []byte type, so we should priority using
// bytes packge to avoid overhead

func getBytes1(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return []byte(sanitize1(string(b))), nil
}

func sanitize1(s string) string {
	return strings.TrimSpace(s)
}

func getBytes2(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize2(b), nil
}

func sanitize2(b []byte) []byte {
	return bytes.TrimSpace(b)
}
