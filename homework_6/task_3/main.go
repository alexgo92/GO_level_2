package main

import (
	"log"
	"sync"
)

func main() {
	counterNow := 0

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(c *int) {
			defer wg.Done()
			*c++
		}(&counterNow)
	}

	wg.Wait()
	log.Println(counterNow)
}
