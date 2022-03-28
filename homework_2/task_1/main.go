package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic caught: %s\n", err)
		}
	}()
	fmt.Println("Please, enter a three-digit number:")

	var num int
	if _, err := fmt.Scanln(&num); err != nil {
		log.Fatalf("error: %s", err)
	}

	if num < 100 || num > 999 {
		log.Fatalf("error: %s", errors.New("integer must have three digits"))
	}

	num1 := num / 100
	num2 := num / 10 % 10
	num3 := num % 10

	if num1 == num2 || num2 == num3 || num1 == num3 {
		log.Fatal("error: three digits must be different")
	}

	if num1 == 0 || num2 == 0 || num3 == 0 {
		log.Fatal("error: three digits must be non-zero")
	}

	slice := []int{num1, num2, num3}

	slicePin := GenerateAListOfNumbers(slice)
	sort.Ints(slicePin)
	fmt.Println("Number combinations:")

	for _, val := range slicePin {
		fmt.Println(val)
	}
}

func GenerateAListOfNumbers(slice []int) []int {
	var slicePin []int

	for i := 0; i < 3; i++ {
		genNum1 := slice[i]*100 + slice[(i+1)%3]*10 + slice[(i+2)%3]
		getNum2 := slice[i]*100 + slice[(i+2)%3]*10 + slice[(i+1)%3]
		slicePin = append(slicePin, genNum1, getNum2)
	}

	return slicePin
}
