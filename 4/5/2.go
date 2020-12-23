/*
Задача 2. Три числа

Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/

package main

import (
	"fmt"
)

func main() {

	var var1, var2, var3 int

	fmt.Println("*Поиск чисел больше 5*")
	fmt.Print("Введите число 1: ")
	fmt.Scan(&var1)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&var2)
	fmt.Print("Введите число 3: ")
	fmt.Scan(&var3)

	if var1 > 5 {
		fmt.Printf("%d > 5\n", var1)
	}

	if var2 > 5 {
		fmt.Printf("%d > 5\n", var2)
	}

	if var3 > 5 {
		fmt.Printf("%d > 5\n", var3)
	}
}
