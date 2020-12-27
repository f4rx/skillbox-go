package main

import (
	"fmt"
)

func main() {
	lineCount := 0
	fmt.Print("Введите размер строк: ")
	fmt.Scan(&lineCount)

	maxLineSize := lineCount*2 - 1

	for i := 1; i <= lineCount; i++ {
		starsCount := i*2 - 1
		spaceCount := (maxLineSize - starsCount) / 2
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < starsCount; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
