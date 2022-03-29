package main

import (
	"fmt"
	"time"
)

/*4.Выполните задание из блока “Для самостоятельного изучения” данной методички
package main

import (
  "fmt"
  "time"
)

func main() {
  defer func() {
     if v := recover(); v != nil {
        fmt.Println("recovered", v)
     }
  }()

  go func() {
     panic("A-A-A!!!")
  }()
  time.Sleep(time.Second)
}

Для самостоятельного изучения. Предложите реализацию примера так, чтобы аварийная остановка программы не выполнилась.*/

func main() {
	// 1 способ - убирем горутину

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	panic("A-A-A!!!")

	// 2 способ - переносим "ловца паники" в саму горутину

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("recovered", v)
			}
		}()
		panic("A-A-A!!!")
	}()
	time.Sleep(2 * time.Second)

	// 3 способ - Нет, не надо слов, не надо паники() =)

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()

	go func() {
		fmt.Println("Лолита - На Титанике")
	}()
	time.Sleep(time.Second)
}
