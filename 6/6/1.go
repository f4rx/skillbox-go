/*
Задание 1. Написание простого цикла
Напишите программу, которая будет принимать от пользователя произвольное число и в цикле выводить на экран все значения от нуля до указанного числа.
*/

package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Введите число")
	fmt.Scan(&input)

	for i := 0; i < input; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
}
