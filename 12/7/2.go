package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const fileName = "lesson2.log"

func main() {
	lineCounter := 1

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	fmt.Printf("Вводите строки, когда надоест напечатайте \"выход\"\n")

	// https://stackoverflow.com/questions/34647039/how-to-use-fmt-scanln-read-from-a-string-separated-by-spaces
	// Не нашёл другого способа как читать строчку из консоли целиком, scanln не работает
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputLine := scanner.Text()
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		if inputLine == "выход" {
			break
		}
		/*
			Вроде про это не рассказывали.... Но как задание, которое заставляет думать норм - возможно было бы уместно в звёздочку
			These are predefined layouts for use in Time.Format and time.Parse. The reference time used in the layouts is the specific time:

			Mon Jan 2 15:04:05 MST 2006
		*/
		fmt.Fprintf(w, "%d %s %s\n", lineCounter, time.Now().Format("2006-01-02 15:04:05"), inputLine)
		w.Flush()
		lineCounter++
	}

}
