package main

import (
	// "bufio"
	"fmt"
	"os"
)

const fileName = "./lesson4.log"

func main() {

	// Удаляем файл, если он есть
	if err := os.Remove(fileName); err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := file.Chmod(0400); err != nil {
		panic(err)
	}

	// Если мы не переоткрываем файл - то можем записывать в файл, т.к. в уже открытом дескрипторе права не меняются
	file.Close()

	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Собственно ошибка, но при открытие файла, а не записи! Err: %v\n", err)
		return
		// panic(err)
	}
	// Надо ли еще раз запускать после закрытия и открытия файла?
	defer file.Close()

	fmt.Printf("Запись в файл\n")
	fmt.Fprintln(file, "ololo")
}
