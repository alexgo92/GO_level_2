package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

/*2.Дополните функцию из п.1 возвратом собственной ошибки в случае возникновения панической ситуации.
Собственная ошибка должна хранить время обнаружения панической ситуации. Критерием успешного выполнения
задания является наличие обработки созданной ошибки в функции main и вывод ее состояния в консоль*/

type JsonStruct struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Price     int    `json:"price"`
	Hero      string `json:"hero"`
}

// создание собственной ошибки
type MyErr struct {
	err     string
	timeNow string
}

func New(err string, time string) error {
	return &MyErr{
		err:     err,
		timeNow: time,
	}
}

func (e MyErr) Error() string {
	return fmt.Sprintf("error: %s\n time: %v\n", e.err, e.timeNow)
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
	// и выводящий собственную ошибку с указанием времени панической ситуации
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			tNow := time.Now()
			formatTimeNow := tNow.Format("15:04")
			strRerecoverErr := fmt.Sprintf("%v", recoverErr)
			err := fmt.Errorf("panic caught: %w", New(strRerecoverErr, formatTimeNow))
			fmt.Println(err)
		}
	}()
	// неявная паника при попытке обратиться к элементу по индексу к batman
	PrintByte(jsonapp.Hero[6])
}

// функция печати байт
func PrintByte(b byte) {
	fmt.Println(b)
}
