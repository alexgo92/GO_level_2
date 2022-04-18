package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/alexgo92/GO_level_2/homework_7/task_2/count"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	// запускаем функцию hello 5 раз
	for i := 0; i < 5; i++ {
		go hello(&wg)
	}
	wg.Wait()

	// передаем в функцию CountFunc имя файла и имя функции
	// и ожидаем что она вернет count = 5
	// поскольку hello() запускалась 5 раз
	count, err := count.CountFunc("main.go", "hello")
	if err != nil {
		log.Fatalf("error: ", err)
	}
	fmt.Println("hello function ran", count, "times")
}

func hello(wg *sync.WaitGroup) {
	fmt.Println("Hi!")
	wg.Done()
}
