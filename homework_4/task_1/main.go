package main

import (
	"fmt"
)

const (
	numberOfGoroutines = 1000
)

func main() {
	ch := make(chan int, 1000)

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			ch <- 1
		}()
	}

	var total int

	for i := 0; i < numberOfGoroutines; i++ {
		<-ch
		total++
	}
	fmt.Println(total)
}
