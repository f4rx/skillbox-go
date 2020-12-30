package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "file.tmp"

func main() {

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	if err := os.Chmod(fileName, 0444); err != nil {
		panic(err)
	}

	w := bufio.NewWriter(file)
	defer file.Close()

	fmt.Printf("Запись в файл\n")

	w.WriteString("ololo")
	w.Flush()
}
