package createFile

import (
	"errors"
	"fmt"
	"os"
)

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(errors.New("error: can't create"), fileName)
	}

}
