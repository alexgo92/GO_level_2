package main

import (
	"fmt"
	"time"
)

func main() {
	var num int

	workers := make(chan struct{}, 1000)

	for i := 0; i < 1000; i++ {
		workers <- struct{}{}

		go func(num *int) {
			defer func() {
				<-workers
			}()
			*num++
			// time.Sleep(time.Second) из за меняющегося итогового результата(num) убрал
		}(&num)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(num)
}
