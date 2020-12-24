/*
Задание 2. Сумма двух чисел по единице
Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом: берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
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

	result := a
	for result < (a + b) {
		result++
		fmt.Printf("%d ", result)
	}
	fmt.Printf("\n")
	fmt.Printf("%d + %d = %d\n", a, b, result)

}
