package main

import (
	"fmt"
)

var (
	v1 = 1
	v2 = 2
	v3 = 3
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	n2 := f3(f2(f1(n)))
	fmt.Println(n2)
}

func f1(n int) int {
	return n + v1
}

func f2(n int) int {
	return n + v2
}

func f3(n int) int {
	return n + v3
}
