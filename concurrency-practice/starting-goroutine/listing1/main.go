package main

func main() {
	newWatcher()

	// Run the application
}

func newWatcher() {
	w := watcher{}
	// we have no way to close this after the main routine called this function
	// exited => goroutines leak
	go w.watch()
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}
