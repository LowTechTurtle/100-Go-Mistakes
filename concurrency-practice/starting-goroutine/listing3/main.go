package main

func main() {
	w := newWatcher()
	// we should just call the close method on the parent routines instead of
	// passing the context, this will block the parent routines until
	// the resources are released
	defer w.close()

	// Run the application
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}

func (w watcher) close() {
	// Close the resources
}
