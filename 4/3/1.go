/*
Задача 1. Минимум

Напишите программу, которая ищет минимум из двух чисел.

Подсказка: учтите, что числа могут быть равны!
*/

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Введите число a: ")
	fmt.Scan(&a)
	fmt.Print("Введите число b: ")
	fmt.Scan(&b)

	if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a == b")
	}
}
