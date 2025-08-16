package main

import (
	"log"
	"os"
)

func listing1(filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func() {
		// should close before the GC collect it
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v\n", err)
		}
	}()

	return nil
}

func writeToFile1(filename string, content []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func() {
		// another benefit is to return the error of closing the file
		// because write is asynchronous, we might close while its writing => error
		// this doesnt guarantee the write on disk
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = f.Write(content)
	return
}

func writeToFile2(filename string, content []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	// we could avoid the error from closing the file and return the error
	// from flushing and commiting the write instead
	return f.Sync()
}
