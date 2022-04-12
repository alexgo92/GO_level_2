package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

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
	wg.Add(10)

	go printLog("world", &wg)

	printLog("beautiful", &wg)
	wg.Wait()
}

func printLog(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		log.Println(s)
		wg.Done()
	}
}
