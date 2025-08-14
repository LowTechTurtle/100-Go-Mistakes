package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

func handler1(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		i := i
		circle := circle
		go func() {
			defer wg.Done()

			result, err := foo(ctx, circle)
			if err != nil {
			}
			// ?
			results[i] = result
		}()
	}

	wg.Wait()
	// ...
	return results, nil
}

func handler2(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)
	// if any call fail and an error return, we will stop all the call
	// and return the first error we got

	for i, circle := range circles {
		i := i
		circle := circle
		// use the custom Go function of err group
		// it take a func error, when the first non nil error is returned
		// it stop all the routines and return that error
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	// wait until the routines finish
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func foo(context.Context, Circle) (Result, error) {
	return Result{}, nil
}

type (
	Circle struct{}
	Result struct{}
)
