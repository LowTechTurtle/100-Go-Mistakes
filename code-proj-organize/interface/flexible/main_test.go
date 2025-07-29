package main

import (
	"bytes"
	"strings"
	"testing"
)

// using small interface can make writing test easier( can create io.Reader and io.Writer)
// using string and bytes, if we used os.File, that would be harder to test
func TestCopy(t *testing.T) {
	const input = "foo"
	source := strings.NewReader(input)
	var b []byte = make([]byte, 0)
	dest := bytes.NewBuffer(b)

	if err := CopySourceDest(source, dest); err != nil {
		t.FailNow()
	}

	if input != dest.String() {
		t.Errorf("expected: %s, got: %s", input, dest.String())
	}
}
