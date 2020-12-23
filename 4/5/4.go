/*
Задача 4. Три числа: еще попытка

Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
*/

package main

import (
	"fmt"
)

func main() {

	var input int
	total := 0

	fmt.Println("Введите первое число")
	fmt.Scan(&input)
	if input > 5 {
		total++
	}

	fmt.Println("Введите второе число")
	fmt.Scan(&input)
	if input > 5 {
		total++
	}

	fmt.Println("Введите третье число")
	fmt.Scan(&input)
	if input > 5 {
		total++
	}

	fmt.Printf("Чисел больше 5: %d\n", total)
}
