package createFile

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(errors.New("error: can't create"), fileName)
	}
	defer file.Close()
}
