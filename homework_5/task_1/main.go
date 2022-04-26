package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var (
		n  int
		wg sync.WaitGroup
	)

	fmt.Print("Enter the number of gorutines to start: ")

	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("error: %s", err)
	}

	// в канале специально указал буфер на 1 единицу больше
	// а вдруг больше запустится горутин (понятно что так неможет быть, но всё-таки...)
	ch := make(chan struct{}, n+1)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			ch <- struct{}{}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	close(ch)

	var countCh int
	for range ch {
		countCh++
	}

	if n == countCh {
		fmt.Println("The program launched", n, "gorutines and waited for them all to complete")
	} else {
		fmt.Println("The program isn't working properly")
	}
}
