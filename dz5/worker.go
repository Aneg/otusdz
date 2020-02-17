package dz5

import (
	"errors"
	"sync"
)

func Run(tasks []func() error, n, m int) error {
	if m <= 0 {
		m = 1
	}
	var err error
	taskCount := len(tasks)
	taskChan := make(chan func() error, taskCount)
	errChan := make(chan error, taskCount)
	completedTaskChan := make(chan bool, taskCount)
	closeChan := make(chan bool)
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go runHandler(&wg, taskChan, completedTaskChan, errChan, closeChan)
	}

	for i := range tasks {
		taskChan <- tasks[i]
	}

	countErrors := 0
	completedTasks := 0
	for {
		select {
		case <-errChan:
			countErrors++
		case <-completedTaskChan:
			completedTasks++
		}

		if countErrors >= m || completedTasks >= taskCount {
			break
		}
	}

	close(taskChan)
	close(closeChan)
	wg.Wait()

	if countErrors >= m {
		err = errors.New("error limit completed")
	}

	close(errChan)
	close(completedTaskChan)

	return err
}

func runHandler(wg *sync.WaitGroup, taskChan chan func() error, completedTaskChan chan bool, errChan chan error, closeChan chan bool) {
	for true {
		select {
		case <-closeChan:
			wg.Done()
			return
		default:
		}
		select {
		case task, ok := <-taskChan:
			if !ok {
				wg.Done()
				return
			}
			if err := task(); err != nil {
				errChan <- err
			}
			completedTaskChan <- true
		}
	}
}
