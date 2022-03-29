package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	timeNow time.Time
}

func New(err string, time time.Time) error {
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
	defer file.Close()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// создаем структуру и считываем в нее даные из json файла
	jsonapp := JsonStruct{}
	err = json.NewDecoder(file).Decode(&jsonapp)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// отложенный вызов обрабатывающий паническую ситуацию
	// и выводящий собственную ошибку с указанием времени панической ситуации
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			tNow := time.Now().UTC()
			rus, err := time.LoadLocation("Europe/Moscow")
			if err != nil {
				fmt.Printf("can't tell the time: %s\n", err)
			}
			formatTimeNow := tNow.In(rus)
			strRerecoverErr := fmt.Sprintf("%v", recoverErr)
			err = fmt.Errorf("panic caught: %w", New(strRerecoverErr, formatTimeNow))
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
