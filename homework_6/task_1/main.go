package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

type MyCount struct {
	sync.Mutex
	counter int
}

func New() *MyCount {
	return &MyCount{counter: 0}
}

func main() {
	fileTrace, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer fileTrace.Close()

	err = trace.Start(fileTrace)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	newCount := New()
	const count = 100

	wg.Add(count)

	for i := 0; i < count; i += 1 {
		go func(newCount *MyCount) {
			defer wg.Done()
			newCount.Lock()
			newCount.counter += 1
			newCount.Unlock()
		}(newCount)
	}
	wg.Wait()
	fmt.Println(newCount.counter)
}
