package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 && n == 0 {
		return nil
	}

	var errCnt int32
	var wg sync.WaitGroup
	taskCh := make(chan Task, len(tasks))

	// Start n workers in goroutines.
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for task := range taskCh {
				if isErrorLimitExceeded(&errCnt, int32(m)) {
					return
				}
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	// Send tasks to the workers.
	for i := 0; i < len(tasks); i++ {
		if isErrorLimitExceeded(&errCnt, int32(m)) {
			break
		}
		taskCh <- tasks[i]
	}

	close(taskCh)
	wg.Wait()

	if errCnt < int32(m) {
		return nil
	}
	return ErrErrorsLimitExceeded
}

// isErrorLimitExceeded checks if the current error count exceeds the limit.
func isErrorLimitExceeded(current *int32, limit int32) bool {
	return atomic.LoadInt32(current) >= limit
}
