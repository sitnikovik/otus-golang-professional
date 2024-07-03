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

	wg.Add(len(tasks))
	limitCh := make(chan struct{}, n)
	stopCh := make(chan struct{})

	for i := 0; i < len(tasks); i++ {
		limitCh <- struct{}{}
		go func(fn func() error) {
			defer func() {
				<-limitCh
				wg.Done()
			}()

			select {
			case <-stopCh:
				return
			default:
			}

			if atomic.LoadInt32(&errCnt) >= int32(m) {
				stopCh <- struct{}{}
				return
			}

			if err := fn(); err != nil {
				atomic.AddInt32(&errCnt, 1)
			}
		}(tasks[i])
	}

	close(limitCh)
	wg.Wait()

	if errCnt < int32(m) {
		return nil
	}
	return ErrErrorsLimitExceeded
}
