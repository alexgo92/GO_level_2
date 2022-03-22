package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

/*1.Для закрепления навыков отложенного вызова функций, напишите программу, содержащую вызов функции, которая будет создавать
паническую ситуацию неявно. Затем создайте отложенный вызов, который будет обрабатывать эту паническую ситуацию и, в частности,
печатать предупреждение в консоль. Критерием успешного выполнения задания является то, что программа не завершается аварийно ни
при каких условиях.*/

type JsonStruct struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Price     int    `json:"price"`
	Hero      string `json:"hero"`
}

func main() {
	file, err := os.Open("./test_data/j.json")
	if err != nil {
		fmt.Println(errors.New("error: file cann't be opened"))
		os.Exit(0)
	}
	// отложенно закрываем файл
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Printf("very bad: %s", err)
		}
	}()

	// создаем структуру и считываем в нее даные из json файла
	jsonapp := JsonStruct{}
	err = json.NewDecoder(file).Decode(&jsonapp)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(0)
	}

	// отложенный вызов обрабатывающий паническую ситуацию
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic caught: %s\n", err)
		}
	}()

	// неявная паника при попытке обратиться к элементу по индексу к batman
	// в принципе, зачем я нагромоздил код выше с json, для меня загадка =)
	PrintByte(jsonapp.Hero[6])
}

// функция печати байт
func PrintByte(b byte) {
	fmt.Println(b)
}
