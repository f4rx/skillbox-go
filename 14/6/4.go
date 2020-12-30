package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	f1(&n)
	f2(&n)

	fmt.Println(n)
}

func f1(n *int) {
	*n = *n * *n
}

func f2(n *int) {
	*n = *n + *n
}
