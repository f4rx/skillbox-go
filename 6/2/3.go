package main

import (
	"fmt"
)

func main(){
	n := 0
	pow := 0

	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Print("Введите степень: ")
	fmt.Scan(&pow)


	result := n
	if pow == 0 && n >= 0 {
		result = 1
	} else if pow == 0 && n < 0 {
		result = -1
	} else {
		for i:=1; i < pow; i++ {
			result *= n
		}
	}

	fmt.Printf("%d\n", result)
}
