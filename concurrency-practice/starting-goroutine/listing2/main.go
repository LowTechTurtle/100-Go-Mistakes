package main

import "context"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// give the watcher the context to cancel when the main routine exit
	newWatcher(ctx)
	// but we only singal that it should be closed, we didnt stop the parent
	// routines to wait for the child rountines clear its resource

	// Run the application
}

func newWatcher(ctx context.Context) {
	w := watcher{}
	// assume watch is context aware
	go w.watch(ctx)
}

type watcher struct { /* Some resources */
}

func (w watcher) watch(context.Context) {}
