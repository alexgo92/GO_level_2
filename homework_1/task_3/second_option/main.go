package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	createFile "github.com/alexgo92/GO_level_2/homework_1/task_3/second_option/create"
)

/*3.Для закрепления практических навыков программирования, напишите программу, которая создаёт
один миллион пустых файлов в известной, пустой директории файловой системы используя вызов os.Create.
Ввиду наличия определенных ограничений операционной системы на число открытых файлов, такая программа
должна выполнять аварийную остановку. Запустите программу и дождитесь полученной ошибки. Используя
отложенный вызов функции закрытия файла, стабилизируйте работу приложения. Критерием успешного выполнения
программы является успешное создание миллиона пустых файлов в директории*/

// указываем директорию куда будут помещаться файлы
var dirPath = flag.String("dirPath", "./test", "folder for files")

func main() {
	flag.Parse()

	// убираем лишние символы директории для функции MkdirAll
	directoryPath := strings.TrimSpace(*dirPath)
	directoryPathWithout := strings.Replace(directoryPath, "./", "", 1)
	// создаем директорию
	err := os.MkdirAll(directoryPathWithout, 0777)
	if err != nil {
		fmt.Printf("can't create a directory: %v\n", err)
		os.Exit(0)
	}

	// создание 100 файлов, а не 1 млн. поскольку система тупит долго, но ошибку все равно не выдает
	// можно просто поменять число 100 на 1 млн.
	for i := 0; i < 100; i++ {
		// второй вариант решения - переносим все (создание и defer) в отдельную функцию CreateFile()
		// но по фату это, наверное, аналогично первому варианту
		fileName := *dirPath + "/" + "file" + strconv.Itoa(i)
		createFile.CreateFile(fileName)
	}
}
