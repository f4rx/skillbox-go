package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	r1 := f1(n)
	r2 := f2(n)

	fmt.Println(n)
	fmt.Println(r1)
	fmt.Println(r2)
}

func f1(n int) (result int) {
	result = n * n
	return
}

func f2(n int) (result int) {
	result = n + n
	return
}
