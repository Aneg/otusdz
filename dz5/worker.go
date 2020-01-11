package dz5

import "time"

var taskChan chan func() error
var errChan chan error
var closeChan chan bool

func Run(tasks []func() error, n, m int) {
	taskChan = make(chan func() error, 10000)
	errChan = make(chan error, 10000)
	closeChan = make(chan bool)
	for i := 0; i < n; i++ {
		go func() {
			for true {
				// можно сделать через атомик. Меньше кода будет)
				select {
				case <-closeChan:
					return
				default:
				}

				select {
				case task := <-taskChan:
					if err := task(); err != nil {
						errChan <- err
					}
				default:
					time.Sleep(time.Second)
				}
			}
		}()
	}

}
