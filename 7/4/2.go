package main

import (
	"fmt"
)

func main() {
	boardSize := 0
	fmt.Print("Введите размер доски: ")
	fmt.Scan(&boardSize)

	var mapChars [2]string
	mapChars[0] = "-"
	mapChars[1] = "*"

	for i := 0; i < boardSize+2; i++ {
		fmt.Print("_")
	}
	fmt.Printf("\n")

	for i := 0; i < boardSize; i++ {
		fmt.Printf("|")
		printChar := i % 2
		for j := 0; j < boardSize; j++ {
			c := mapChars[(printChar+j)%2]
			fmt.Print(c)
		}
		fmt.Printf("|\n")
	}

	for i := 0; i < boardSize+2; i++ {
		fmt.Print("⁻")
	}
	fmt.Printf("\n")

}
