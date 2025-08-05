package main

import "os"

func readFiles1(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		// if the function readFiles1 never return, then these defer never runs
		// => leak memory
		defer file.Close()
		// now do something with file
	}
	return nil
}

func readFiles2(ch <-chan string) error {
	for path := range ch {
		// we can fix this by calling a readFile function
		// every file will call one function and if it executed normally then
		// all defer up until the error will run => no mem leak
		err := readFiles(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func readFiles(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// do sth with file
	return nil
}

func readFiles3(ch <-chan string) error {
	for path := range ch {
		// we also can do the same trick with readFile2 with a closure
		// but writing a stand alone function is more readable and easier to
		// write unit test if needed
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			//do sth with file
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
