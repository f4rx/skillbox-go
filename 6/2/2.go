package main

import (
	"fmt"
)

func main(){

	result := 0

	for i:=0; i<4; i++ {
		fmt.Println("Введите число")
		input := 0
		fmt.Scan(&input)
		result += input
	}

	fmt.Printf("%d\n", result)
}
