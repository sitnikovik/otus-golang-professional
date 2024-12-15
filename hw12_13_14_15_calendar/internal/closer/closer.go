package closer

import (
	"context"
	"sync"
)

// closerFunc is a function that closes a resource.
type closerFunc = func(ctx context.Context) error

var (
	// mu is a mutex for the closers slice
	mu sync.Mutex
	// closers is a slice of closer functions
	closers []closerFunc
)

// Add adds a closer function to the list of closers.
func Add(closer closerFunc) {
	mu.Lock()
	defer mu.Unlock()
	closers = append(closers, closer)
}

// CloseAll closes all closers and returns the first error occurred.
func CloseAll(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()

	var wg sync.WaitGroup
	var err error

	for _, closer := range closers {
		wg.Add(1)
		go func(c closerFunc) {
			defer wg.Done()
			if e := c(ctx); e != nil {
				err = e
			}
		}(closer)
	}

	wg.Wait()
	return err
}
