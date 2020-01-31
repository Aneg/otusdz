package dz5

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

func Run(tasks []func() error, n, m int) error {
	if m <= 0 {
		m = 1
	}
	var err error
	taskCount := uint32(len(tasks))
	taskChan := make(chan func() error, taskCount)
	errChan := make(chan error, taskCount)
	closeChan := make(chan bool)
	wg := sync.WaitGroup{}

	var completedHandlerCount uint32 = 0
	wg.Add(n)
	for i := 0; i < n; i++ {
		go runHandler(&wg, taskChan, errChan, closeChan, &completedHandlerCount)
	}

	for i := range tasks {
		taskChan <- tasks[i]
	}

	countErrors := 0
	ticker := time.NewTicker(500 * time.Millisecond)
	for _ = range ticker.C {
		if !(countErrors < m && atomic.LoadUint32(&completedHandlerCount) < taskCount) {
			break
		}
		select {
		case <-errChan:
			countErrors++
		default:
		}
	}
	ticker.Stop()

	if countErrors >= m {
		err = errors.New("error limit completed")
	}
	close(closeChan)
	wg.Wait()
	return err
}

func runHandler(wg *sync.WaitGroup, taskChan chan func() error, errChan chan error, closeChan chan bool, completedCount *uint32) {
	for true {
		// можно сделать через атомик. Меньше кода будет)
		select {
		case <-closeChan:
			wg.Done()
			return
		default:
		}

		select {
		case task := <-taskChan:
			if err := task(); err != nil {
				errChan <- err
			}
			atomic.AddUint32(completedCount, 1)
		default:
			// TODO: через тики в for
			time.Sleep(time.Second)
		}
	}
}
