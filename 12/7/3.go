package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const fileName = "lesson2.log"

func main() {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ваариант 1
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	message1 := make([]byte, fileStat.Size())

	if _, err = file.Read(message1); err != nil {
		panic(err)
	}

	fmt.Printf("Размер файла %s _%d_ байт (из file.Stat()). Прочитан целиком file.Read(message1) \n", fileName, fileStat.Size())
	fmt.Printf("%s\n", message1)

	file.Seek(0, 0) // Возвращаем позицию в 0
	// Вариант когда не знаем размер файла
	message2 := make([]byte, 0)
	counter := 0
	for {
		buf := make([]byte, 10) // у меня на тесте файл 48 байт, 10 оптимально для теста, в реальности надо килобайта 4 наверно делать
		n, err := file.Read(buf)

		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		if n == 0 {
			break
		}
		// fmt.Println(b)
		counter += n
		message2 = append(message2, buf...)
	}
	fmt.Printf("Размер файла %s _%d_ байт при побайтовом чтение (file.Read()))\n", fileName, counter)
	fmt.Printf("%s\n", message2)

	// Вариант 2 с io/ioutil (Задание 6.2)
	fileAll, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Размер файла %s _%d_ байт при чтение файла целиком в память с ioutil.ReadFile()\n", fileName, len(fileAll))
	fmt.Printf("%s\n", fileAll)
}

/*
Заметки на память
http://zetcode.com/golang/readfile/
https://kgrz.io/reading-files-in-go-an-overview.html
*/
