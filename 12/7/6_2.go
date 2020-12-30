package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const fileName = "lesson6-2.log"

func main() {
	lineCounter := 1
	message := ""

	fmt.Printf("Вводите строки, когда надоест напечатайте \"выход\"\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		if inputLine == "выход" {
			break
		}

		message = message + fmt.Sprintf("%d %s %s\n", lineCounter, time.Now().Format("2006-01-02 15:04:05"), inputLine)

		lineCounter++
	}

	fmt.Printf("Перед записью в буфере (памяти) накопилось %d байт\n", len([]byte(message)))
	if err := ioutil.WriteFile(fileName, []byte(message), 0644); err != nil {
		panic(err)
	}

}
