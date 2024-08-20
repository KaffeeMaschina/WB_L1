package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	val int
	mu  sync.Mutex
}

func main() {
	withMutex()
	withWaitGroup()
}
func withWaitGroup() {
	var wg sync.WaitGroup
	newCounter := counter{}

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			newCounter.val++
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(newCounter.val)
}
func withMutex() {
	newCounter := counter{}
	for i := 0; i <= 5; i++ {
		go func() {
			newCounter.mu.Lock()
			newCounter.val++
			newCounter.mu.Unlock()
		}()

	}
	time.Sleep(1 * time.Second)
	fmt.Println(newCounter.val)
}
