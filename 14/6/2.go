package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println(isEven(n))

}

func isEven (n int) bool {
	return n%2 == 0
}
