package main

import "fmt"

func main() {
	fmt.Println("Введите число a для расчёта квадрата числа")

	var i int
	fmt.Scan(&i)
	fmt.Printf("%d^2 = %d\n", i, i * i)
}
