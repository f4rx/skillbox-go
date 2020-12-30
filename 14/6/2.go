package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println(evenOdd(n))

}

func evenOdd(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
