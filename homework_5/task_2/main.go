package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	mu    sync.Mutex
	count int
}

func main() {
	var (
		counter MutexCounter
		wg      sync.WaitGroup
	)

	for i := 0; i < 100000; i++ {
		wg.Add(1)

		go func(c *MutexCounter) {
			c.mu.Lock()
			defer c.mu.Unlock()
			c.count++
			wg.Done()
		}(&counter)
	}
	wg.Wait()
	fmt.Println("Програма завершилась успешно, количество запущеных горутин: ", counter.count)
}
