package main

import "io"

// since io.Reader and io.Writer is a small interface, it is
// powerful in the sense of abstraction and reusability
func CopySourceDest(source io.Reader, dest io.Writer) error {
	b, err := io.ReadAll(source)
	if err != nil {
		return err
	}

	_, err = dest.Write(b)
	return err
}
