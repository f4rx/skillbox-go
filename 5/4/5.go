/*
5. Определение максимальных процентов

Что нужно сделать

Задача учебная, и человек с ней справится сам, но давайте научим нашу программу определять, какие две из предложенных ставок по депозитам являются максимальными. Пользователь будет вводить три процентные ставки, а программа должна вывести, какие две из них являются более выгодными.

Советы и рекомендации

Задачу можно решать несколькими способами, например, от противного.
*/

package main

import (
	"fmt"
)

func main() {
	var rate1, rate2, rate3 int

	fmt.Println("Поиск выгодных ставок")

	fmt.Print("Введите ставку 1: ")
	fmt.Scan(&rate1)
	fmt.Print("Введите ставку 2: ")
	fmt.Scan(&rate2)
	fmt.Print("Введите ставку 3: ")
	fmt.Scan(&rate3)

	if rate1 < rate2 && rate1 < rate3 {
		fmt.Printf("2 выгодные ставки: %d и %d\n", rate2, rate3)
	} else if rate2 < rate1 && rate2 < rate3 {
		fmt.Printf("2 выгодные ставки: %d и %d\n", rate1, rate3)
	} else {
		fmt.Printf("2 выгодные ставки: %d и %d\n", rate1, rate2)
	}

}