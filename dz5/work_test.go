package dz5

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tests := make([]func() error, 0, 20)

	var errorCount int32 = 0

	for i := 0; i < 20; i++ {
		tests = append(tests, func() error {
			randSecond := rand.Int() % 5 * 1000 * 1000 * 1000
			time.Sleep(time.Duration(randSecond))
			atomic.AddInt32(&errorCount, 1)
			return fmt.Errorf("error")
		})
	}
	if err := Run(tests, 4, 10); err == nil {
		t.Errorf("Отсутствует ошибка о достижении лиммита ошибок")
	}

	if atomic.LoadInt32(&errorCount) < 10 {
		t.Errorf("Отработало не верное количество функций")
	}
}
