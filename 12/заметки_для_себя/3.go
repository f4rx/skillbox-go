package main

import (
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

const fileName = "lesson2.log"

func main() {

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ваариант 1, простой
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Размер файла %s _%d_ байт (из file.Stat())\n", fileName, fileStat.Size())

	// Ваариант 2 с побайтовым чтением
	counter := 0
	for {
		b := make([]byte, 10) // у меня на тесте файл 48 байт, 10 оптимально для теста, в реальности надо килобайта 4 наверно делать
		n, err := file.Read(b)

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
	}
	fmt.Printf("Размер файла %s _%d_ байт при побайтовом чтение (file.Read()))\n", fileName, counter)

	// Вариант 3 с io/ioutil (Задание 6.2)
	fileAll, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Размер файла %s _%d_ байт при чтение файла целиком в память с ioutil.ReadFile()\n", fileName, len(fileAll))
}

/*
Заметки на память
http://zetcode.com/golang/readfile/
https://kgrz.io/reading-files-in-go-an-overview.html
*/
