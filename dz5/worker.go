package dz5

import (
	"sync"
	"sync/atomic"
	"time"
)

func Run(tasks []func() error, n, m int) error {

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

	go func(tasks []func() error) {
		for i := range tasks {
			taskChan <- tasks[i]
		}
	}(tasks)

	countErrors := 0
	for countErrors < m && atomic.LoadUint32(&completedHandlerCount) < taskCount {
		select {
		case <-errChan:
			countErrors++
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
	closeChan <- true
	wg.Wait()
	return nil
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
