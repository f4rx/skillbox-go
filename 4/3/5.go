/*
Задача 5. Кратность числа

Напишите программу, которая проверяет, делится ли одно число на другое без остатка.*/

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

	if a % b == 0 {
		fmt.Println("A делится на B без остатка")
	} else {
		fmt.Println("Остаток от деления A на B: ", a % b )
	}

}
